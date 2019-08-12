package obj

import (
	"github.com/zhang1career/lib/ast"
)

type Obj struct {
}

func New() ast.Calculable {
	return &Obj{}
}

func (this *Obj) Calc(params []ast.Valuable) interface{} {
	return params[0].Evaluate()
}