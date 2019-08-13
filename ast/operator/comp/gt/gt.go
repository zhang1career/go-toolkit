package gt

import "github.com/zhang1career/lib/ast"

type Gt struct {
}

func New() ast.Calculable {
	return &Gt{}
}

func (this *Gt) Calc(params []ast.Valuable) interface{} {
	return params[0].Evaluate().(int) > params[1].Evaluate().(int)
}