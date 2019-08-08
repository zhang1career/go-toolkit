package operator

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/ast/operator/sql/ins"
	"github.com/zhang1career/lib/ast/operator/sql/sel"
)

var OperatorMap = map[string]interface{} {
	"SELECT": sel.New,
	"INSERT": ins.New,
}

func New(op string) (ast.Calculable, error) {
	return sel.New(), nil
}