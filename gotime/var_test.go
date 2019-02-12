package gotime

import "testing"

var arr = []int{1, 2, 3}
var str = "hello, world"

func TestVarType(t *testing.T) {
	t.Logf("arr's type: %s", VarType(arr))
	t.Logf("str's type: %s", VarType(str))
}
