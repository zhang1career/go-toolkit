package charray

import (
	"github.com/zhang1career/lib/ast"
)

type Charray struct {
	value string
}

func New(val interface{}) ast.Valuable {
	return &Charray{val.(string)}
}

func (this *Charray) GetValue() interface{} {
	return this.value
}