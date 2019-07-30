package restapi

import "database/sql/driver"

// NewConnector returns new driver.Connector.
func NewConnector(cfg *Config) (driver.Connector, error) {
	cfg = cfg.Clone()
	// normalize the contents of cfg so calls to NewConnector have the same
	// behavior as MySQLDriver.OpenConnector
	if err := cfg.normalize(); err != nil {
		return nil, err
	}
	return &connector{cfg: cfg}, nil
}

// OpenConnector implements driver.DriverContext.
func (d restConn) OpenConnector(dsn string) (driver.Connector, error) {
	cfg, err := ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	return &connector{
		cfg: cfg,
	}, nil
}
