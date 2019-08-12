package cond

import (
	"fmt"
	"github.com/zhang1career/lib/ast"
)

type Cond struct {
}

func New() ast.Calculable {
	return &Cond{}
}

func (this *Cond) Calc(params []ast.Valuable) interface{} {
	url := fmt.Sprintf("%s", params[0].Evaluate())
	return url
}