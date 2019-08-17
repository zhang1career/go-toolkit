package gt

import "github.com/zhang1career/lib/programmer/ast"

type Gt struct {
	value string
}

func New() ast.Calculable {
	return &Gt{">"}
}

func (this *Gt) Calc(params []ast.Evaluable) interface{} {
	return params[0].Evaluate().(int) > params[1].Evaluate().(int)
}

func (this *Gt) GetValue() string {
	return this.value
}