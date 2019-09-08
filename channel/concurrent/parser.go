package concurrent

import (
	"github.com/zhang1career/lib/channel/ctrlbus"
)

func CreateParser() *parser {
	return &parser{
		managers:   make([]*ctrlbus.Ctrlbus, 0),
		seekers:    make(chan chan interface{}),
		input:      make(chan interface{}),
		output:     make(chan interface{}),
	}
}

func (p *parser) AddTeam(config map[string]interface{}, num int, f func(i int) Work) {
	if num <= 0 {
		return
	}

	start := len(p.managers)
	for i := start; i < start + num; i++ {
		manager := ctrlbus.CreateCtrlbus(config)
		p.managers = append(p.managers, manager)

		team := createWorker(i, f(i))
		team.run(manager, p.seekers, p.output)
	}
}

func (p *parser) Run() {
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


func (p *parser) Parse(in interface{}) interface{} {
	p.input <- in
	output := <-p.output
	return output
}