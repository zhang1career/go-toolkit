package gotime

import "reflect"

func VarType (v interface{}) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}
