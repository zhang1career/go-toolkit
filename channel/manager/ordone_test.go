package manager_test

import (
	"github.com/zhang1career/lib/channel/manager"
	"testing"
)

func TestOrDone(t *testing.T) {
	done := make(chan interface{})
	
	d1 := make(chan interface{})
	d2 := make(chan interface{})
	d3 := make(chan interface{})
	d4 := make(chan interface{})
	
	
	out, err := manager.OrDone(done, d1, d2, d3, d4)
	if err != nil {
		t.Log(err.Error())
		return
	}
	
	close(d2)
	
	for v := range out {
		t.Log(v)
	}
	
}