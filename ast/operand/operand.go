package operand

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/ast/operand/instance/integer"
)

func New(param interface{}) ast.Valuable{
	return integer.New(param)
}
