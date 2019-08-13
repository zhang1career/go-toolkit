package operator

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/ast/operator/calc/add"
	"github.com/zhang1career/lib/ast/operator/comp/gt"
	"github.com/zhang1career/lib/ast/operator/sql/cond"
	"github.com/zhang1career/lib/ast/operator/sql/obj"
	"github.com/zhang1career/lib/ast/operator/sql/sel"
	"github.com/zhang1career/lib/gotime"
)

var OperatorMap = map[string]interface{} {
	// calc
	"+":        add.New,
	// compare
	">":        gt.New,
	// sql
	"select":   sel.New,
	"target":   obj.New,
	"from":     obj.New,
	"where":    cond.New,
}

func New(op string) (ast.Calculable, error) {
	ret, err := gotime.Call(OperatorMap, op)
	return ret[0].Interface().(ast.Calculable), err
}