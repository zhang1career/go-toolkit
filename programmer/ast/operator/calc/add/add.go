package add

import (
	"github.com/zhang1career/lib/programmer/ast"
)

type Add struct {
	value string
}

func New() ast.Calculable {
	return &Add{"add"}
}

func (this *Add) Calc(params []ast.Evaluable) interface{} {
	ret := 0
	for _, param := range params {
		ret += param.Evaluate().(int)
	}
	return ret
}

func (this *Add) GetValue() string {
	return this.value
}