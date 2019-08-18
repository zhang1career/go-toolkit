package gotime

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)


func Locate(skip ...int) (fun string, file string, line int, err error) {
	sk := 1
	if len(skip) > 0 && skip[0] > 1 {
		sk = skip[0]
	}
	
	pc, file, line, ok := runtime.Caller(sk)
	if !ok {
		return "", "", 0, fmt.Errorf("runtime.Caller failed")
	}
	
	fun = runtime.FuncForPC(pc).Name()
	ix := strings.LastIndex(fun, ".")
	if ix > 0 && (ix+1) < len(fun) {
		fun = fun[ix+1:]
	}
	
	nd, nf := filepath.Split(file)
	file = filepath.Join(filepath.Base(nd), nf)
	
	return fun, file, line, nil
}

func WhichFunc(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func ThisFunc() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func Call(i interface{}, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(i)
	if len(params) != f.Type().NumIn() {
		return nil, fmt.Errorf("call with wrong number of params")
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in), nil
}

func CallFromMap(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, fmt.Errorf("call %s with wrong number of params", name)
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in), nil
}
