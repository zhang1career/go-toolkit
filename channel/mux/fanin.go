package mux

import (
	"fmt"
	"github.com/zhang1career/lib/channel/manager"
)

func Fanin(m *manager.Manager, varIns ...<-chan interface{}) (<-chan interface{}, error) {
	if len(varIns) <= 0 {
		return nil, fmt.Errorf("no input given")
	}
	
	out := make(chan interface{})
	
	m.CreateWaitGroup("ins")
	m.WaitAdd("ins", len(varIns))
	for _, in := range varIns {
		transmit(m, out, in)
	}
	
	go func() {
		m.Wait("ins")
		close(out)
	}()
	
	return out, nil
}

func transmit(m *manager.Manager, out chan<- interface{}, in <-chan interface{}) {
	go func() {
		defer m.WaitDone("ins")
		for i := range in {
			select {
			case <-m.Done:
				return
			case out <- i:
			}
		}
	}()
}
