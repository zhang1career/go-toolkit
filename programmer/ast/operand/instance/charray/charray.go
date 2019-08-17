package charray

import (
	"github.com/zhang1career/lib/programmer/ast"
	"github.com/zhang1career/lib/log"
)

type Charray struct {
	value string
}

func New(val interface{}) ast.Evaluable {
	return &Charray{val.(string)}
}

func (this *Charray) Evaluate() interface{} {
	log.Trace("%s", this.value)
	return this.value
}

func (this *Charray) GetValue() string {
	return this.value
}