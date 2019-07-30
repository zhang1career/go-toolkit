package restapi

import (
	"database/sql/driver"
	"io"
	"math"
	"reflect"
)

type resultSet struct {
	columns     []restField
	columnNames []string
	done        bool
}

type restRows struct {
	rc     *restConn
	rs     resultSet
	finish func()
}

type binaryRows struct {
	restRows
}

type textRows struct {
	restRows
}

func (rows *restRows) Columns() []string {
	if rows.rs.columnNames != nil {
		return rows.rs.columnNames
	}

	columns := make([]string, len(rows.rs.columns))
	if rows.rc != nil && rows.rc.cfg.ColumnsWithAlias {
		for i := range columns {
			if tableName := rows.rs.columns[i].tableName; len(tableName) > 0 {
				columns[i] = tableName + "." + rows.rs.columns[i].name
			} else {
				columns[i] = rows.rs.columns[i].name
			}
		}
	} else {
		for i := range columns {
			columns[i] = rows.rs.columns[i].name
		}
	}

	rows.rs.columnNames = columns
	return columns
}

func (rows *restRows) ColumnTypeDatabaseTypeName(i int) string {
	return rows.rs.columns[i].typeDatabaseName()
}

// func (rows *restRows) ColumnTypeLength(i int) (length int64, ok bool) {
// 	return int64(rows.rs.columns[i].length), true
// }

func (rows *restRows) ColumnTypeNullable(i int) (nullable, ok bool) {
	return rows.rs.columns[i].flags&flagNotNULL == 0, true
}

func (rows *restRows) ColumnTypePrecisionScale(i int) (int64, int64, bool) {
	column := rows.rs.columns[i]
	decimals := int64(column.decimals)

	switch column.fieldType {
	case fieldTypeDecimal, fieldTypeNewDecimal:
		if decimals > 0 {
			return int64(column.length) - 2, decimals, true
		}
		return int64(column.length) - 1, decimals, true
	case fieldTypeTimestamp, fieldTypeDateTime, fieldTypeTime:
		return decimals, decimals, true
	case fieldTypeFloat, fieldTypeDouble:
		if decimals == 0x1f {
			return math.MaxInt64, math.MaxInt64, true
		}
		return math.MaxInt64, decimals, true
	}

	return 0, 0, false
}

func (rows *restRows) ColumnTypeScanType(i int) reflect.Type {
	return rows.rs.columns[i].scanType()
}

func (rows *restRows) Close() (err error) {
	if f := rows.finish; f != nil {
		f()
		rows.finish = nil
	}

	mc := rows.rc
	if mc == nil {
		return nil
	}
	if err := mc.error(); err != nil {
		return err
	}

	// flip the buffer for this connection if we need to drain it.
	// note that for a successful query (i.e. one where rows.next()
	// has been called until it returns false), `rows.rc` will be nil
	// by the time the user calls `(*Rows).Close`, so we won't reach this
	// see: https://github.com/golang/go/commit/651ddbdb5056ded455f47f9c494c67b389622a47
	mc.buf.flip()

	// Remove unread packets from stream
	if !rows.rs.done {
		err = mc.readUntilEOF()
	}
	if err == nil {
		if err = mc.discardResults(); err != nil {
			return err
		}
	}

	rows.rc = nil
	return err
}

func (rows *restRows) HasNextResultSet() (b bool) {
	if rows.rc == nil {
		return false
	}
	return rows.rc.status&statusMoreResultsExists != 0
}

func (rows *restRows) nextResultSet() (int, error) {
	if rows.rc == nil {
		return 0, io.EOF
	}
	if err := rows.rc.error(); err != nil {
		return 0, err
	}

	// Remove unread packets from stream
	if !rows.rs.done {
		if err := rows.rc.readUntilEOF(); err != nil {
			return 0, err
		}
		rows.rs.done = true
	}

	if !rows.HasNextResultSet() {
		rows.rc = nil
		return 0, io.EOF
	}
	rows.rs = resultSet{}
	return rows.rc.readResultSetHeaderPacket()
}

func (rows *restRows) nextNotEmptyResultSet() (int, error) {
	for {
		resLen, err := rows.nextResultSet()
		if err != nil {
			return 0, err
		}

		if resLen > 0 {
			return resLen, nil
		}

		rows.rs.done = true
	}
}

func (rows *binaryRows) NextResultSet() error {
	resLen, err := rows.nextNotEmptyResultSet()
	if err != nil {
		return err
	}

	rows.rs.columns, err = rows.rc.readColumns(resLen)
	return err
}

func (rows *binaryRows) Next(dest []driver.Value) error {
	if mc := rows.rc; mc != nil {
		if err := mc.error(); err != nil {
			return err
		}

		// Fetch next row from stream
		return rows.readRow(dest)
	}
	return io.EOF
}

func (rows *textRows) NextResultSet() (err error) {
	resLen, err := rows.nextNotEmptyResultSet()
	if err != nil {
		return err
	}

	rows.rs.columns, err = rows.rc.readColumns(resLen)
	return err
}

func (rows *textRows) Next(dest []driver.Value) error {
	if mc := rows.rc; mc != nil {
		if err := mc.error(); err != nil {
			return err
		}

		// Fetch next row from stream
		return rows.readRow(dest)
	}
	return io.EOF
}
