package operator

import (
	"github.com/zhang1career/golab/gotime"
	"github.com/zhang1career/golab/compiler/ast"
	"github.com/zhang1career/golab/compiler/ast/operator/calc/add"
	"github.com/zhang1career/golab/compiler/ast/operator/comp/gt"
	"github.com/zhang1career/golab/compiler/ast/operator/sql/cond"
	"github.com/zhang1career/golab/compiler/ast/operator/sql/from"
	"github.com/zhang1career/golab/compiler/ast/operator/sql/sel"
	"github.com/zhang1career/golab/compiler/ast/operator/sql/target"
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
	ret, err := gotime.CallFromMap(OperatorMap, op)
	return ret[0].Interface().(ast.Calculable), err
}