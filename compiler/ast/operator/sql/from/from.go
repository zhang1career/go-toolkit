package from

import (
	"github.com/zhang1career/lib/compiler/ast"
)

type From struct {
	value string
}

func New() ast.Calculable {
	return &From{"from"}
}

func (this *From) Calc(params []ast.Evaluable) interface{} {
	return params[0].Evaluate()
}

func (this *From) GetValue() string {
	return this.value
}