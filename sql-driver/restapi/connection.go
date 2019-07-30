package restapi

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

type restConn struct {
	buf              buffer
	netConn          net.Conn
	rawConn          net.Conn // underlying connection when netConn is TLS connection.
	affectedRows     uint64
	insertId         uint64
	cfg              *Config
	maxAllowedPacket int
	maxWriteSize     int
	writeTimeout     time.Duration
	flags            clientFlag
	status           statusFlag
	sequence         uint8
	parseTime        bool
	reset            bool // set when the Go SQL package calls ResetSession

	// for context support (Go 1.8+)
	watching bool
	watcher  chan<- context.Context
	closech  chan struct{}
	finished chan<- struct{}
	canceled atomicError // set non-nil if conn is canceled
	closed   atomicBool  // set when conn is closed, before closech is closed
}

// Handles parameters set in DSN after the connection is established
func (rc *restConn) handleParams() (err error) {
	for param, val := range rc.cfg.Params {
		switch param {
		// Charset
		case "charset":
			charsets := strings.Split(val, ",")
			for i := range charsets {
				// ignore errors here - a charset may not exist
				err = rc.exec("SET NAMES " + charsets[i])
				if err == nil {
					break
				}
			}
			if err != nil {
				return
			}

		// System Vars
		default:
			err = rc.exec("SET " + param + "=" + val + "")
			if err != nil {
				return
			}
		}
	}

	return
}

func (rc *restConn) markBadConn(err error) error {
	if rc == nil {
		return err
	}
	if err != errBadConnNoWrite {
		return err
	}
	return driver.ErrBadConn
}

func (rc *restConn) Begin() (driver.Tx, error) {
	return rc.begin(false)
}

func (rc *restConn) begin(readOnly bool) (driver.Tx, error) {
	if rc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	var q string
	if readOnly {
		q = "START TRANSACTION READ ONLY"
	} else {
		q = "START TRANSACTION"
	}
	err := rc.exec(q)
	if err == nil {
		return &restTx{rc}, err
	}
	return nil, rc.markBadConn(err)
}

func (rc *restConn) Close() (err error) {
	// Makes Close idempotent
	if !rc.closed.IsSet() {
		err = rc.writeCommandPacket(comQuit)
	}

	rc.cleanup()

	return
}

// Closes the network connection and unsets internal variables. Do not call this
// function after successfully authentication, call Close instead. This function
// is called before auth or on auth failure because MySQL will have already
// closed the network connection.
func (rc *restConn) cleanup() {
	if !rc.closed.TrySet(true) {
		return
	}

	// Makes cleanup idempotent
	close(rc.closech)
	if rc.netConn == nil {
		return
	}
	if err := rc.netConn.Close(); err != nil {
		errLog.Print(err)
	}
}

func (rc *restConn) error() error {
	if rc.closed.IsSet() {
		if err := rc.canceled.Value(); err != nil {
			return err
		}
		return ErrInvalidConn
	}
	return nil
}

func (rc *restConn) Prepare(query string) (driver.Stmt, error) {
	if rc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	// Send command
	err := rc.writeCommandPacketStr(comStmtPrepare, query)
	if err != nil {
		return nil, rc.markBadConn(err)
	}

	stmt := &restStmt{
		rc: rc,
	}

	// Read Result
	columnCount, err := stmt.readPrepareResultPacket()
	if err == nil {
		if stmt.paramCount > 0 {
			if err = rc.readUntilEOF(); err != nil {
				return nil, err
			}
		}

		if columnCount > 0 {
			err = rc.readUntilEOF()
		}
	}

	return stmt, err
}

func (rc *restConn) interpolateParams(query string, args []driver.Value) (string, error) {
	// Number of ? should be same to len(args)
	if strings.Count(query, "?") != len(args) {
		return "", driver.ErrSkip
	}

	buf, err := rc.buf.takeCompleteBuffer()
	if err != nil {
		// can not take the buffer. Something must be wrong with the connection
		errLog.Print(err)
		return "", ErrInvalidConn
	}
	buf = buf[:0]
	argPos := 0

	for i := 0; i < len(query); i++ {
		q := strings.IndexByte(query[i:], '?')
		if q == -1 {
			buf = append(buf, query[i:]...)
			break
		}
		buf = append(buf, query[i:i+q]...)
		i += q

		arg := args[argPos]
		argPos++

		if arg == nil {
			buf = append(buf, "NULL"...)
			continue
		}

		switch v := arg.(type) {
		case int64:
			buf = strconv.AppendInt(buf, v, 10)
		case uint64:
			// Handle uint64 explicitly because our custom ConvertValue emits unsigned values
			buf = strconv.AppendUint(buf, v, 10)
		case float64:
			buf = strconv.AppendFloat(buf, v, 'g', -1, 64)
		case bool:
			if v {
				buf = append(buf, '1')
			} else {
				buf = append(buf, '0')
			}
		case time.Time:
			if v.IsZero() {
				buf = append(buf, "'0000-00-00'"...)
			} else {
				v := v.In(rc.cfg.Loc)
				v = v.Add(time.Nanosecond * 500) // To round under microsecond
				year := v.Year()
				year100 := year / 100
				year1 := year % 100
				month := v.Month()
				day := v.Day()
				hour := v.Hour()
				minute := v.Minute()
				second := v.Second()
				micro := v.Nanosecond() / 1000

				buf = append(buf, []byte{
					'\'',
					digits10[year100], digits01[year100],
					digits10[year1], digits01[year1],
					'-',
					digits10[month], digits01[month],
					'-',
					digits10[day], digits01[day],
					' ',
					digits10[hour], digits01[hour],
					':',
					digits10[minute], digits01[minute],
					':',
					digits10[second], digits01[second],
				}...)

				if micro != 0 {
					micro10000 := micro / 10000
					micro100 := micro / 100 % 100
					micro1 := micro % 100
					buf = append(buf, []byte{
						'.',
						digits10[micro10000], digits01[micro10000],
						digits10[micro100], digits01[micro100],
						digits10[micro1], digits01[micro1],
					}...)
				}
				buf = append(buf, '\'')
			}
		case []byte:
			if v == nil {
				buf = append(buf, "NULL"...)
			} else {
				buf = append(buf, "_binary'"...)
				if rc.status&statusNoBackslashEscapes == 0 {
					buf = escapeBytesBackslash(buf, v)
				} else {
					buf = escapeBytesQuotes(buf, v)
				}
				buf = append(buf, '\'')
			}
		case string:
			buf = append(buf, '\'')
			if rc.status&statusNoBackslashEscapes == 0 {
				buf = escapeStringBackslash(buf, v)
			} else {
				buf = escapeStringQuotes(buf, v)
			}
			buf = append(buf, '\'')
		default:
			return "", driver.ErrSkip
		}

		if len(buf)+4 > rc.maxAllowedPacket {
			return "", driver.ErrSkip
		}
	}
	if argPos != len(args) {
		return "", driver.ErrSkip
	}
	return string(buf), nil
}

func (rc *restConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	if rc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	if len(args) != 0 {
		if !rc.cfg.InterpolateParams {
			return nil, driver.ErrSkip
		}
		// try to interpolate the parameters to save extra roundtrips for preparing and closing a statement
		prepared, err := rc.interpolateParams(query, args)
		if err != nil {
			return nil, err
		}
		query = prepared
	}
	rc.affectedRows = 0
	rc.insertId = 0

	err := rc.exec(query)
	if err == nil {
		return &restResult{
			affectedRows: int64(rc.affectedRows),
			insertId:     int64(rc.insertId),
		}, err
	}
	return nil, rc.markBadConn(err)
}

// Internal function to execute commands
func (rc *restConn) exec(query string) error {
	// Send command
	if err := rc.writeCommandPacketStr(comQuery, query); err != nil {
		return rc.markBadConn(err)
	}

	// Read Result
	resLen, err := rc.readResultSetHeaderPacket()
	if err != nil {
		return err
	}

	if resLen > 0 {
		// columns
		if err := rc.readUntilEOF(); err != nil {
			return err
		}

		// rows
		if err := rc.readUntilEOF(); err != nil {
			return err
		}
	}

	return rc.discardResults()
}

func (rc *restConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	return rc.query(query, args)
}

func (rc *restConn) query(query string, args []driver.Value) (*textRows, error) {
	if rc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	if len(args) != 0 {
		if !rc.cfg.InterpolateParams {
			return nil, driver.ErrSkip
		}
		// try client-side prepare to reduce roundtrip
		prepared, err := rc.interpolateParams(query, args)
		if err != nil {
			return nil, err
		}
		query = prepared
	}
	// Send command
	err := rc.writeCommandPacketStr(comQuery, query)
	if err == nil {
		// Read Result
		var resLen int
		resLen, err = rc.readResultSetHeaderPacket()
		if err == nil {
			rows := new(textRows)
			rows.rc = rc

			if resLen == 0 {
				rows.rs.done = true

				switch err := rows.NextResultSet(); err {
				case nil, io.EOF:
					return rows, nil
				default:
					return nil, err
				}
			}

			// Columns
			rows.rs.columns, err = rc.readColumns(resLen)
			return rows, err
		}
	}
	return nil, rc.markBadConn(err)
}

// Gets the value of the given MySQL System Variable
// The returned byte slice is only valid until the next read
func (rc *restConn) getSystemVar(name string) ([]byte, error) {
	// Send command
	if err := rc.writeCommandPacketStr(comQuery, "SELECT @@"+name); err != nil {
		return nil, err
	}

	// Read Result
	resLen, err := rc.readResultSetHeaderPacket()
	if err == nil {
		rows := new(textRows)
		rows.rc = rc
		rows.rs.columns = []restField{{fieldType: fieldTypeVarChar}}

		if resLen > 0 {
			// Columns
			if err := rc.readUntilEOF(); err != nil {
				return nil, err
			}
		}

		dest := make([]driver.Value, resLen)
		if err = rows.readRow(dest); err == nil {
			return dest[0].([]byte), rc.readUntilEOF()
		}
	}
	return nil, err
}

// finish is called when the query has canceled.
func (rc *restConn) cancel(err error) {
	rc.canceled.Set(err)
	rc.cleanup()
}

// finish is called when the query has succeeded.
func (rc *restConn) finish() {
	if !rc.watching || rc.finished == nil {
		return
	}
	select {
	case rc.finished <- struct{}{}:
		rc.watching = false
	case <-rc.closech:
	}
}

// Ping implements driver.Pinger interface
func (rc *restConn) Ping(ctx context.Context) (err error) {
	if rc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return driver.ErrBadConn
	}

	if err = rc.watchCancel(ctx); err != nil {
		return
	}
	defer rc.finish()

	if err = rc.writeCommandPacket(comPing); err != nil {
		return rc.markBadConn(err)
	}

	return rc.readResultOK()
}

// BeginTx implements driver.ConnBeginTx interface
func (rc *restConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if err := rc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer rc.finish()

	if sql.IsolationLevel(opts.Isolation) != sql.LevelDefault {
		level, err := mapIsolationLevel(opts.Isolation)
		if err != nil {
			return nil, err
		}
		err = rc.exec("SET TRANSACTION ISOLATION LEVEL " + level)
		if err != nil {
			return nil, err
		}
	}

	return rc.begin(opts.ReadOnly)
}

func (rc *restConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := rc.watchCancel(ctx); err != nil {
		return nil, err
	}

	rows, err := rc.query(query, dargs)
	if err != nil {
		rc.finish()
		return nil, err
	}
	rows.finish = rc.finish
	return rows, err
}

func (rc *restConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := rc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer rc.finish()

	return rc.Exec(query, dargs)
}

func (rc *restConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if err := rc.watchCancel(ctx); err != nil {
		return nil, err
	}

	stmt, err := rc.Prepare(query)
	rc.finish()
	if err != nil {
		return nil, err
	}

	select {
	default:
	case <-ctx.Done():
		stmt.Close()
		return nil, ctx.Err()
	}
	return stmt, nil
}

func (stmt *restStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := stmt.rc.watchCancel(ctx); err != nil {
		return nil, err
	}

	rows, err := stmt.query(dargs)
	if err != nil {
		stmt.rc.finish()
		return nil, err
	}
	rows.finish = stmt.rc.finish
	return rows, err
}

func (stmt *restStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := stmt.rc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer stmt.rc.finish()

	return stmt.Exec(dargs)
}

func (rc *restConn) watchCancel(ctx context.Context) error {
	if rc.watching {
		// Reach here if canceled,
		// so the connection is already invalid
		rc.cleanup()
		return nil
	}
	// When ctx is already cancelled, don't watch it.
	if err := ctx.Err(); err != nil {
		return err
	}
	// When ctx is not cancellable, don't watch it.
	if ctx.Done() == nil {
		return nil
	}
	// When watcher is not alive, can't watch it.
	if rc.watcher == nil {
		return nil
	}

	rc.watching = true
	rc.watcher <- ctx
	return nil
}

func (rc *restConn) startWatcher() {
	watcher := make(chan context.Context, 1)
	rc.watcher = watcher
	finished := make(chan struct{})
	rc.finished = finished
	go func() {
		for {
			var ctx context.Context
			select {
			case ctx = <-watcher:
			case <-rc.closech:
				return
			}

			select {
			case <-ctx.Done():
				rc.cancel(ctx.Err())
			case <-finished:
			case <-rc.closech:
				return
			}
		}
	}()
}

func (rc *restConn) CheckNamedValue(nv *driver.NamedValue) (err error) {
	nv.Value, err = converter{}.ConvertValue(nv.Value)
	return
}

// ResetSession implements driver.SessionResetter.
// (From Go 1.10)
func (rc *restConn) ResetSession(ctx context.Context) error {
	if rc.closed.IsSet() {
		return driver.ErrBadConn
	}
	rc.reset = true
	return nil
}
