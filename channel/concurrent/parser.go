package concurrent

import (
	"github.com/zhang1career/lib/channel/ctrlbus"
)

type Parser struct {
	Output      chan interface{}
	manager     []*ctrlbus.Ctrlbus
	seekers     chan chan interface{}
}

func CreateParser() *Parser {
	return &Parser{
		Output:  make(chan interface{}),
		manager: make([]*ctrlbus.Ctrlbus, 0),
		seekers: make(chan chan interface{}),
	}
}

func (this *Parser) AddWorker(num int, works <-chan Work, workerIds chan<- int) {
	if num <= 0 {
		return
	}
	
	for i := 0; i < num; i++ {
		ctrl := ctrlbus.CreateCtrlbus()
		this.manager = append(this.manager, ctrl)
		
		workerIds <- i
		work := <-works
		worker := createWorker(i, work)
		worker.run(ctrl, this.seekers, this.Output)
	}
}

func (this *Parser) Run(jobs chan interface{}) {
	go func() {
		var seekerQ = make([]chan interface{}, 0)
		var jobQ = make([]interface{}, 0)
		
		for {
			var seeker chan interface{}
			var job interface{}
			
			if len(seekerQ) > 0 && len(jobQ) > 0 {
				seeker = seekerQ[0]
				job = jobQ[0]
			}
			
			select {
			case s := <-this.seekers:
				seekerQ = append(seekerQ, s)
			case j := <-jobs:
				jobQ = append(jobQ, j)
			case seeker <- job:
				seekerQ = seekerQ[1:]
				jobQ = jobQ[1:]
			}
		}
	}()
}
