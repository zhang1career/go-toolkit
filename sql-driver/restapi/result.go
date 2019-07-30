package restapi

type restResult struct {
	affectedRows int64
	insertId     int64
}

func (res *restResult) LastInsertId() (int64, error) {
	return res.insertId, nil
}

func (res *restResult) RowsAffected() (int64, error) {
	return res.affectedRows, nil
}
