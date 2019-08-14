package integer

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/log"
)

type Integer struct {
	value int
}

func New(val interface{}) ast.Valuable {
	return &Integer{val.(int)}
}

func (this *Integer) Evaluate() interface{} {
	log.Info("%d", this.value)
	return this.value
}