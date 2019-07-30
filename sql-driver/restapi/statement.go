package restapi

import (
	"database/sql/driver"
	"fmt"
	"io"
	"reflect"
)

type restStmt struct {
	rc         *restConn
	id         uint32
	paramCount int
}

func (stmt *restStmt) Close() error {
	if stmt.rc == nil || stmt.rc.closed.IsSet() {
		// driver.Stmt.Close can be called more than once, thus this function
		// has to be idempotent.
		// See also Issue #450 and golang/go#16019.
		//errLog.Print(ErrInvalidConn)
		return driver.ErrBadConn
	}

	err := stmt.rc.writeCommandPacketUint32(comStmtClose, stmt.id)
	stmt.rc = nil
	return err
}

func (stmt *restStmt) NumInput() int {
	return stmt.paramCount
}

func (stmt *restStmt) ColumnConverter(idx int) driver.ValueConverter {
	return converter{}
}

func (stmt *restStmt) Exec(args []driver.Value) (driver.Result, error) {
	if stmt.rc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	// Send command
	err := stmt.writeExecutePacket(args)
	if err != nil {
		return nil, stmt.rc.markBadConn(err)
	}

	rc := stmt.rc

	rc.affectedRows = 0
	rc.insertId = 0

	// Read Result
	resLen, err := rc.readResultSetHeaderPacket()
	if err != nil {
		return nil, err
	}

	if resLen > 0 {
		// Columns
		if err = rc.readUntilEOF(); err != nil {
			return nil, err
		}

		// Rows
		if err := rc.readUntilEOF(); err != nil {
			return nil, err
		}
	}

	if err := rc.discardResults(); err != nil {
		return nil, err
	}

	return &restResult{
		affectedRows: int64(rc.affectedRows),
		insertId:     int64(rc.insertId),
	}, nil
}

func (stmt *restStmt) Query(args []driver.Value) (driver.Rows, error) {
	return stmt.query(args)
}

func (stmt *restStmt) query(args []driver.Value) (*binaryRows, error) {
	if stmt.rc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	// Send command
	err := stmt.writeExecutePacket(args)
	if err != nil {
		return nil, stmt.rc.markBadConn(err)
	}

	rc := stmt.rc

	// Read Result
	resLen, err := rc.readResultSetHeaderPacket()
	if err != nil {
		return nil, err
	}

	rows := new(binaryRows)

	if resLen > 0 {
		rows.rc = rc
		rows.rs.columns, err = rc.readColumns(resLen)
	} else {
		rows.rs.done = true

		switch err := rows.NextResultSet(); err {
		case nil, io.EOF:
			return rows, nil
		default:
			return nil, err
		}
	}

	return rows, err
}

type converter struct{}

// ConvertValue mirrors the reference/default converter in database/sql/driver
// with _one_ exception.  We support uint64 with their high bit and the default
// implementation does not.  This function should be kept in sync with
// database/sql/driver defaultConverter.ConvertValue() except for that
// deliberate difference.
func (c converter) ConvertValue(v interface{}) (driver.Value, error) {
	if driver.IsValue(v) {
		return v, nil
	}

	if vr, ok := v.(driver.Valuer); ok {
		sv, err := callValuerValue(vr)
		if err != nil {
			return nil, err
		}
		if !driver.IsValue(sv) {
			return nil, fmt.Errorf("non-Value type %T returned from Value", sv)
		}
		return sv, nil
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr:
		// indirect pointers
		if rv.IsNil() {
			return nil, nil
		} else {
			return c.ConvertValue(rv.Elem().Interface())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return rv.Float(), nil
	case reflect.Bool:
		return rv.Bool(), nil
	case reflect.Slice:
		ek := rv.Type().Elem().Kind()
		if ek == reflect.Uint8 {
			return rv.Bytes(), nil
		}
		return nil, fmt.Errorf("unsupported type %T, a slice of %s", v, ek)
	case reflect.String:
		return rv.String(), nil
	}
	return nil, fmt.Errorf("unsupported type %T, a %s", v, rv.Kind())
}

var valuerReflectType = reflect.TypeOf((*driver.Valuer)(nil)).Elem()

// callValuerValue returns vr.Value(), with one exception:
// If vr.Value is an auto-generated method on a pointer type and the
// pointer is nil, it would panic at runtime in the panicwrap
// method. Treat it like nil instead.
//
// This is so people can implement driver.Value on value types and
// still use nil pointers to those types to mean nil/NULL, just like
// string/*string.
//
// This is an exact copy of the same-named unexported function from the
// database/sql package.
func callValuerValue(vr driver.Valuer) (v driver.Value, err error) {
	if rv := reflect.ValueOf(vr); rv.Kind() == reflect.Ptr &&
		rv.IsNil() &&
		rv.Type().Elem().Implements(valuerReflectType) {
		return nil, nil
	}
	return vr.Value()
}
