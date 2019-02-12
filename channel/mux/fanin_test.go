package mux_test

import (
	"github.com/zhang1career/lib/channel/manager"
	"github.com/zhang1career/lib/channel/mux"
	"github.com/zhang1career/lib/channel/source"
	"testing"
)

func TestFanin(t *testing.T) {
	m := manager.CreateManager()
	
	in1 := source.VariadicSource(1, 2, 3, 4, 5)
	in2 := source.VariadicSource("a", "b", "c", "d", "e", "f", "g")
	in3 := source.RandSource(6)
	
	out, err := mux.Fanin(m, in1, in2, in3)
	if err != nil {
		t.Log(err.Error())
		return
	}
	
	for v := range out {
		t.Log(v)
	}
	
	m.Destroy()
}