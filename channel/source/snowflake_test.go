package source

import (
	"testing"
)

func TestSnowflaker_Run(t *testing.T) {
	worker := CreateWorker(1)
	done := make(chan interface{})
	
	out := worker.Run(done)
	for {
		v, ok := <-out
		if !ok {
			t.Log("done")
			return
		}
		
		result := worker.Decompose(v)
		t.Logf("%v", result.t)
		t.Logf("%x", result.m)
		t.Logf("%x", result.s)
	}
}