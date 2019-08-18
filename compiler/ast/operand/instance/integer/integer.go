package integer

import (
	"github.com/zhang1career/lib/compiler/ast"
	"github.com/zhang1career/lib/log"
	"strconv"
)

type Integer struct {
	value int
}

func New(val interface{}) ast.Evaluable {
	return &Integer{val.(int)}
}

func (this *Integer) Evaluate() interface{} {
	log.Trace("%d", this.value)
	return this.value
}

func (this *Integer) GetValue() string {
	return strconv.Itoa(this.value)
}