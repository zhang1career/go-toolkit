package gotime

import (
	"fmt"
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
	_, err := Call(funcs, "foo")
	if err != nil {
		fmt.Println(err.Error())
	}
	
	_, err = Call(funcs, "bar", 1, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
}