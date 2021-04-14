package filter_test

import (
	"github.com/zhang1career/golab/channel/filter"
	"github.com/zhang1career/golab/channel/source"

	"testing"
)

func TestTake(t *testing.T) {
	done := make(chan interface{})
	
	in  := source.Repeat(done, 1,2,3)
	out := filter.Take(done, in, 7)
	
	for i:=0; i<5; i++ {
		v, ok := <-out
		if !ok {
			t.Log("done")
		} else {
			t.Log(v)
		}
	}
	
	close(done)
	//time.Sleep(time.Second)
	
	for i:=0; i<5; i++ {
		v, ok := <-out
		if !ok {
			t.Log("done")
		} else {
			t.Log(v)
		}
	}
}
