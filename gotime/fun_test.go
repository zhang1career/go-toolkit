package gotime_test

import (
	"fmt"
	"github.com/zhang1career/lib/gotime"
	"testing"
)

func foo() {
	fmt.Println("foo")
}
func bar(a, b, c int) {
	fmt.Printf("a=%d, b=%d, c=%d", a, b, c)
}

var funcs = map[string]interface{} {
	"foo": foo,
	"bar": bar,
}

func TestCall(t *testing.T) {
	_, err := gotime.Call(funcs, "foo")
	if err != nil {
		t.Error(err.Error())
	}
	
	_, err = gotime.Call(funcs, "bar", 1, 2, 3)
	if err != nil {
		t.Log(err.Error())
	}
}