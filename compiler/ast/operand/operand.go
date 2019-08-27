package operand

import (
	"github.com/zhang1career/lib/compiler/ast"
	"github.com/zhang1career/lib/compiler/ast/operand/instance/charray"
	"github.com/zhang1career/lib/compiler/ast/operand/instance/integer"
	"github.com/zhang1career/lib/gotime"
	"reflect"
)

func New(param interface{}) ast.Evaluable{
	switch gotime.VarType(param) {
	case reflect.Int:
		return integer.New(param)
	case reflect.String:
		return charray.New(param)
	default:
		return nil
	}
}


