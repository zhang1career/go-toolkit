package charray

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/log"
)

type Charray struct {
	value string
}

func New(val interface{}) ast.Valuable {
	return &Charray{val.(string)}
}

func (this *Charray) Evaluate() interface{} {
	log.Info("%s", this.value)
	return this.value
}