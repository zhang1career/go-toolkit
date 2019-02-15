package source_test

import (
	"fmt"
	"github.com/zhang1career/lib/channel/ctrlbus"
	"github.com/zhang1career/lib/channel/source"
	"testing"
)


type parser struct {
	manager     []*ctrlbus.Ctrlbus
	seekers     chan chan uint64
	Output      chan uint64
}

func createParser() *parser {
	return &parser{
		manager: make([]*ctrlbus.Ctrlbus, 0),
		seekers: make(chan chan uint64),
		Output:  make(chan uint64),
	}
}

func (this *parser) AddWorker(num int) {
	if num <= 0 {
		return
	}
	
	for i := 0; i < num; i++ {
		ctrl := ctrlbus.CreateCtrlbus()
		this.manager = append(this.manager, ctrl)
		
		sf := source.CreateSnowFlake(i)
		worker := createWorker(i, sf)
		worker.run(ctrl, this.seekers, this.Output)
	}
}

func (this *parser) run(jobs chan uint64) {
	go func() {
		var seekerQ = make([]chan uint64, 0)
		var jobQ = make([]uint64, 0)
		
		for {
			var seeker chan uint64
			var job uint64
			
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


type work interface {
	Do(uint64) uint64
}

type worker struct {
	id      int
	work
}

func createWorker(id int, work work) *worker {
	return &worker{id: id, work: work}
}

func (this *worker) run(ctrlbus *ctrlbus.Ctrlbus, seekers chan chan uint64, output chan<- uint64) {
	go func() {
		var seeker = make(chan uint64)
		defer close(seekers)
		for {
			seekers <- seeker
			select {
			case <-ctrlbus.Done:
				return
			case job := <-seeker:
				out := this.Do(job)
				output <- out
				fmt.Println(this.id)
			}
		}
	}()
}


func TestSnowflake_Run(t *testing.T) {
	p := createParser()
	p.AddWorker(3)
	
	salt := make(chan uint64)
	p.run(salt)
	
	salt <- 0
	for i:=0; i<10; i++ {
		v, ok := <-p.Output
		if !ok {
			t.Log("done")
			return
		}
		
		time, machine, serial := source.Decompose(v)
		t.Logf("%v", time)
		t.Logf("%x", machine)
		t.Logf("%x", serial)
		
		salt <- v
	}
}
