package concurrent

import (
	"github.com/zhang1career/lib/channel/ctrlbus"
)

func CreateParser() *Parser {
	return &Parser{
		managers:   make([]*ctrlbus.Ctrlbus, 0),
		seekers:    make(chan chan interface{}),
		input:      make(chan interface{}),
		output:     make(chan Output),
	}
}

func (p *Parser) AddTeam(config map[string]interface{}, num int, f func(*ctrlbus.Ctrlbus, int) Work) {
	if num <= 0 {
		return
	}

	start := len(p.managers)
	for i := start; i < start + num; i++ {
		manager := ctrlbus.CreateCtrlbus(config)
		p.managers = append(p.managers, manager)

		worker := createWorker(i, f(manager, i))
		worker.run(manager, p.seekers, p.output)
	}
}

func (p *Parser) Run() {
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
			case s := <-p.seekers:
				seekerQ = append(seekerQ, s)
			case j := <-p.input:
				jobQ = append(jobQ, j)
			case seeker <- job:
				seekerQ = seekerQ[1:]
				jobQ = jobQ[1:]
			}
		}
	}()
}


func (p *Parser) Parse(in interface{}) Output {
	p.input <- in
	output := <-p.output
	return output
}

func CreateOutput(val interface{}, err error) Output {
	return Output{val,err}
}

func (o Output) GetValue() interface{} {
	return o.value
}

func (o Output) GetErr() error {
	return o.err
}