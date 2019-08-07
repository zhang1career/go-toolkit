package operation

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/ast/operand"
	"github.com/zhang1career/lib/ast/operator/sql/sel"
	"github.com/zhang1career/lib/gotime"
	"github.com/zhang1career/lib/log"
	"reflect"
)

func New(param interface{}) ast.Valuable {
	if gotime.VarType(param) != reflect.Map {
		return operand.New(param)
	}
	
	_, ok := param.(map[string]interface{})
	if !ok {
		log.Fatal("Parameter error. Parameter is not a map, param=%s", param)
	}
	
	var ret ast.Valuable
	//for _, _ := range paramMap {
		//op, err := gotime.Call(operator.OperatorMap, key, value)
		op := sel.New()
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
		ret = reflect.ValueOf(op).Interface().(ast.Valuable)
		//break
	//}
	
	return ret
}
