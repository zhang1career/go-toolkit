package operation

import (
	"github.com/zhang1career/lib/compiler"
	"github.com/zhang1career/lib/compiler/ast"
	"github.com/zhang1career/lib/compiler/ast/operand"
	"github.com/zhang1career/lib/compiler/ast/operator"
	"github.com/zhang1career/lib/gotime"
	"github.com/zhang1career/lib/log"
	"reflect"
)

type Operation struct {
	operator ast.Calculable
	operands []ast.Evaluable
}


func ImportFromMap(param interface{}) ast.Evaluable {
	// operand
	if gotime.VarType(param) != reflect.Map {
		return operand.New(param)
	}
	
	paramMap, ok := param.(compiler.Dim)
	if !ok {
		log.Fatal("Parameter error. Parameter is not a map, param=%s", param)
	}
	
	var op ast.Calculable
	var err error
	var paramArray = make([]ast.Evaluable, 0)
	for k, v := range paramMap {
		op, err = operator.New(k)
		if err != nil {
			log.Fatal(err.Error())
		}
		paramType := gotime.VarType(v)
		// operation
		if paramType == reflect.Map {
			for vk, vv := range v.(compiler.Dim) {
				paramArray = append(paramArray, ImportFromMap(compiler.Dim{vk: vv}))
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

func (this *Operation) ExportToMap() {
}


func ImportFromString(param string) ast.Evaluable {
	return nil
}

func (this *Operation) ExportToString() {
}


func (this *Operation) Evaluate() interface{} {
	log.Trace("%s.Calc", gotime.WhichObj(this.operator))
	return this.operator.Calc(this.operands)
}