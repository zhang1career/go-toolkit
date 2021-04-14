package gate_test

import (
	"github.com/zhang1career/golab/channel/gate"
	"github.com/zhang1career/golab/channel/source"
	"testing"
)

func TestOr(t *testing.T) {
	done := make(chan interface{})
	
	in1 := source.VariadicSource(1, 2, 3, 4, 5)
	in2 := source.VariadicSource("a", "b", "c", "d", "e", "f", "g")
	in3 := source.RandSource(6)
	
	out, err := gate.Or(done, in1, in2, in3)
	if err != nil {
		t.Log(err.Error())
		return
	}
	
	for v := range out {
		t.Log(v)
	}
	close(done)
	
}