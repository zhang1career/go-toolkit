package operation

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/ast/operand"
	"github.com/zhang1career/lib/ast/operator"
	"github.com/zhang1career/lib/gotime"
	"github.com/zhang1career/lib/log"
	"reflect"
)

type Operation struct {
	operator ast.Calculable
	operands []ast.Valuable
}

func New(param interface{}) ast.Valuable {
	// operand
	if gotime.VarType(param) != reflect.Map {
		return operand.New(param)
	}
	
	paramMap, ok := param.(ast.Item)
	if !ok {
		log.Fatal("Parameter error. Parameter is not a map, param=%s", param)
	}
	
	var op ast.Calculable
	var err error
	var paramArray = make([]ast.Valuable, 0)
	for k, v := range paramMap {
		op, err = operator.New(k)
		if err != nil {
			log.Fatal(err.Error())
		}
		paramType := gotime.VarType(v)
		// operation
		if paramType == reflect.Map {
			for vk, vv := range v.(ast.Item) {
				paramArray = append(paramArray, New(ast.Item{vk: vv}))
			}
			break
		}

		// operands
		if paramType == reflect.Slice {
			for _, vi := range v.([]interface{}) {
				paramArray = append(paramArray, operand.New(vi))
			}
	    } else {
			paramArray = append(paramArray, operand.New(v))
		}
		break
	}

	return &Operation{op, paramArray}
}

func (this *Operation) Evaluate() interface{} {
	return this.operator.Calc(this.operands)
}