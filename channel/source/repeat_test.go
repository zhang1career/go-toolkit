package source_test

import (
	"github.com/zhang1career/lib/channel/source"
	"testing"
)

func TestRepeat(t *testing.T) {
	done := make(chan interface{})
	
	out := source.Repeat(done, 1,2,3)
	
	for i:=0; i<10; i++ {
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