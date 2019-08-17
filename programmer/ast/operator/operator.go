package operator

import (
	"github.com/zhang1career/lib/programmer/ast"
	"github.com/zhang1career/lib/programmer/ast/operator/calc/add"
	"github.com/zhang1career/lib/programmer/ast/operator/comp/gt"
	"github.com/zhang1career/lib/programmer/ast/operator/sql/cond"
	"github.com/zhang1career/lib/programmer/ast/operator/sql/from"
	"github.com/zhang1career/lib/programmer/ast/operator/sql/sel"
	"github.com/zhang1career/lib/programmer/ast/operator/sql/target"
	"github.com/zhang1career/lib/gotime"
)

var OperatorMap = map[string]interface{} {
	// calc
	"+":        add.New,
	// compare
	">":        gt.New,
	// sql
	"select":   sel.New,
	"target":   target.New,
	"from":     from.New,
	"where":    cond.New,
}

func New(op string) (ast.Calculable, error) {
	ret, err := gotime.Call(OperatorMap, op)
	return ret[0].Interface().(ast.Calculable), err
}