package sa

type Syntax interface {
	GetValue() string
}

type Operator interface {
	Syntax
	GetPreCount()   int
	GetPostCount()  int
}