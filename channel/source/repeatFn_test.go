package source_test

import (
	"github.com/zhang1career/lib/channel/source"
	"math/rand"
	"testing"
)

func r() interface{} {
	return rand.Int()
}
func TestRepeatFn(t *testing.T) {
	done := make(chan interface{})
	
	out := source.RepeatFn(done, r)
	
	for i:=0; i<5; i++ {
		v, ok := <-out
		if !ok {
			t.Log("done")
		} else {
			t.Log(v)
		}
	}
	
	close(done)
	
	for i:=0; i<10; i++ {
		v, ok := <-out
		if !ok {
			t.Log("done")
		} else {
			t.Log(v)
		}
	}
}

