package cond

import (
	"fmt"
	"github.com/zhang1career/lib/ast"
)

type Cond struct {
	value string
}

func New() ast.Calculable {
	return &Cond{"where"}
}

func (this *Cond) Calc(params []ast.Evaluable) interface{} {
	url := fmt.Sprintf("%s", params[0].Evaluate())
	return url
}

func (this *Cond) GetValue() string {
	return this.value
}