package integer

import "github.com/zhang1career/lib/ast"

type Integer struct {
	value int
}

func New(val interface{}) ast.Valuable {
	return &Integer{val.(int)}
}

func (this *Integer) GetValue() interface{} {
	return this.value
}