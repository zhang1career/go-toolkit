package add

import "github.com/zhang1career/lib/ast"

type Add struct {
}

func New() ast.Calculable {
	return &Add{}
}

func (this *Add) Calc(params []ast.Valuable) interface{} {
	return params[0].Evaluate().(int) + params[1].Evaluate().(int)
}