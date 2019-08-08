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
	operands map[string]ast.Valuable
}

func New(param interface{}) ast.Valuable {
	if gotime.VarType(param) != reflect.Map {
		return operand.New(param)
	}
	
	paramMap, ok := param.(ast.Item)
	if !ok {
		log.Fatal("Parameter error. Parameter is not a map, param=%s", param)
	}
	
	var op ast.Calculable
	var para map[string]ast.Valuable
	var err error
	for k, v := range paramMap {
		op, err = operator.New(k)
		if err != nil {
			log.Fatal(err.Error())
		}
		for vk, vv := range v.(ast.Item) {
			para[vk] = New(vv)
		}
		break
	}
	
	return &Operation{op, nil}
}

func (this *Operation) Evaluate() interface{} {
	return this.operator.Calc(this.operands)
}