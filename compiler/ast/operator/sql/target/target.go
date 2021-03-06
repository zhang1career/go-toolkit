package target

import (
	"github.com/zhang1career/golab/compiler/ast"
)

type Target struct {
	value string
}

func New() ast.Calculable {
	return &Target{"target"}
}

func (this *Target) Calc(params []ast.Evaluable) interface{} {
	return params[0].Evaluate()
}

func (this *Target) GetValue() string {
	return this.value
}