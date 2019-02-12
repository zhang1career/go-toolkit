package back

type Error struct {
	ErrCode int
	ErrMsg  string
	UserMsg string
}

func (e *Error) Code() int {
	return e.ErrCode
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func (e *Error) Message() string {
	return e.UserMsg
}
