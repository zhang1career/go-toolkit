package restapi

import (
	"context"
	"database/sql/driver"
	"github.com/zhang1career/lib/log"
	"net"
)

type connector struct {
	cfg *Config // immutable private copy.
}

// Connect implements driver.Connector interface.
// Connect returns a connection to the database.
func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	var err error

	// New restConn
	rc := &restConn{
		maxAllowedPacket: maxPacketSize,
		maxWriteSize:     maxPacketSize - 1,
		closech:          make(chan struct{}),
		cfg:              c.cfg,
	}
	rc.parseTime = rc.cfg.ParseTime

	// Connect to Server
	dialsLock.RLock()
	dial, ok := dials[rc.cfg.Net]
	dialsLock.RUnlock()
	if ok {
		rc.netConn, err = dial(ctx, rc.cfg.Addr)
	} else {
		nd := net.Dialer{Timeout: rc.cfg.Timeout}
		rc.netConn, err = nd.DialContext(ctx, rc.cfg.Net, rc.cfg.Addr)
	}

	if err != nil {
		if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
			log.Error("net.Error from Dial()': ", nerr.Error())
			return nil, driver.ErrBadConn
		}
		return nil, err
	}

	// Disable TCP Keepalives on TCP connections
	if tc, ok := rc.netConn.(*net.TCPConn); ok {
		if err := tc.SetKeepAlive(false); err != nil {
			// Don't send COM_QUIT before handshake.
			rc.netConn.Close()
			rc.netConn = nil
			return nil, err
		}
	}

	// Call startWatcher for context support (From Go 1.8)
	rc.startWatcher()
	if err := rc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer rc.finish()

	rc.buf = newBuffer(rc.netConn)

	// Set I/O timeouts
	rc.buf.timeout = rc.cfg.ReadTimeout
	rc.writeTimeout = rc.cfg.WriteTimeout

	// Reading Handshake Initialization Packet
	//authData, plugin, err := rc.readHandshakePacket()
	//if err != nil {
	//	rc.cleanup()
	//	return nil, err
	//}
	//
	//if plugin == "" {
	//	plugin = defaultAuthPlugin
	//}

	// Send Client Authentication Packet
	//authResp, err := rc.auth(authData, plugin)
	//if err != nil {
	//	// try the default auth plugin, if using the requested plugin failed
	//	log.Error("could not use requested auth plugin '"+plugin+"': ", err.Error())
	//	plugin = defaultAuthPlugin
	//	authResp, err = rc.auth(authData, plugin)
	//	if err != nil {
	//		rc.cleanup()
	//		return nil, err
	//	}
	//}
	//if err = rc.writeHandshakeResponsePacket(authResp, plugin); err != nil {
	//	rc.cleanup()
	//	return nil, err
	//}

	// Handle response to auth packet, switch methods if possible
	//if err = rc.handleAuthResult(authData, plugin); err != nil {
	//	// Authentication failed and MySQL has already closed the connection
	//	// (https://dev.mysql.com/doc/internals/en/authentication-fails.html).
	//	// Do not send COM_QUIT, just cleanup and return the error.
	//	rc.cleanup()
	//	return nil, err
	//}

	if rc.cfg.MaxAllowedPacket > 0 {
		rc.maxAllowedPacket = rc.cfg.MaxAllowedPacket
	} else {
		// Get max allowed packet size
		maxap, err := rc.getSystemVar("max_allowed_packet")
		if err != nil {
			rc.Close()
			return nil, err
		}
		rc.maxAllowedPacket = stringToInt(maxap) - 1
	}
	if rc.maxAllowedPacket < maxPacketSize {
		rc.maxWriteSize = rc.maxAllowedPacket
	}

	// Handle DSN Params
	err = rc.handleParams()
	if err != nil {
		rc.Close()
		return nil, err
	}

	return rc, nil
}

// Driver implements driver.Connector interface.
// Driver returns &MySQLDriver{}.
func (c *connector) Driver() driver.Driver {
	return &RestDriver{}
}