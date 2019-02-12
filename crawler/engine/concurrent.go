package engine

import (
	"github.com/zhang1career/lib/log"
)

type Scheduler interface {
	Config(map[string]interface{})
	Submit(interface{})
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) start(seed ...Request) {
	for _, s := range seed {
		e.Scheduler.Submit(s)
	}
}

func (e *ConcurrentEngine) createWorker(requests chan interface{}, answers chan interface{}) {
	w := worker{dumps:make(map[string]bool)}
	go func() {
		for request := range requests {
			// mark root
			w.root(request)
			// exit due to condition
			if w.done(request) {
				return
			}
			// explore the tree
			result, err := w.explore(request)
			if err != nil {
				log.Error(err.Error())
				continue
			}
			// parse result
			children, out := w.parse(result)
			for _, child := range children {
				go func() { requests <- child }()
			}
			for _, result := range out.([]interface{}) {
				answers <- result
			}
		}
	}()
}


func (e *ConcurrentEngine) Run(seed ...Request) {
	requests := make(chan interface{})
	answers  := make(chan interface{})
	
	e.Scheduler.Config(map[string]interface{}{
		"request_chan": requests,
	})

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(requests, answers)
	}

	e.start(seed...)
	for {
		for a := range answers {
			log.Info("got item: %v", a)
		}
		for request := range requests {
			e.Scheduler.Submit(request)
		}
	}
}


