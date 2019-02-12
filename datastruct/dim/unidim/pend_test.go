package unidim

import "testing"

func TestPrepend(t *testing.T) {
	var intRet []interface{}

	intARR321:= []interface{}{3, 2, 1}
	intPEND4 := 4
	intPEND5 := 5
	
	intRet = Prepend(intARR321)
	t.Logf("%v", intRet)
	
	intRet = Prepend(intARR321, intPEND4)
	t.Logf("%v", intRet)
	
	intRet = Prepend(intARR321, intPEND4, intPEND5)
	t.Logf("%v", intRet)
	
	
	intARRabc:= []interface{}{"c", 0, "a"}
	intPENDd := "d"
	intPENDe := "雷猴，"
	
	intRet = Prepend(intARRabc, intPENDd, intPENDe)
	t.Logf("%v", intRet)
}