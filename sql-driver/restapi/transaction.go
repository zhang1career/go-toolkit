package restapi

type restTx struct {
	rc *restConn
}

func (tx *restTx) Commit() (err error) {
	if tx.rc == nil || tx.rc.closed.IsSet() {
		return ErrInvalidConn
	}
	err = tx.rc.exec("COMMIT")
	tx.rc = nil
	return
}

func (tx *restTx) Rollback() (err error) {
	if tx.rc == nil || tx.rc.closed.IsSet() {
		return ErrInvalidConn
	}
	err = tx.rc.exec("ROLLBACK")
	tx.rc = nil
	return
}
