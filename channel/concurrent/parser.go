package concurrent

import (
	"github.com/zhang1career/lib/channel/ctrlbus"
)

func CreateParser() *parser {
	return &parser{
		managers:   make([]*ctrlbus.Ctrlbus, 0),
		wholesaler: &wholesaler{},
		tasks:      make(chan *task),
		products:   make(chan *product),
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

		team := createTeam(i, f(i))
		team.run(manager, p.wholesaler, p.products)
	}
}

func (p *parser) Run() {
	go func() {
		var teamQ = make([]*team, 0)
		var taskQ = make([]*task, 0)

		for {
			var tm0 *team
			var tk0 *task

			if len(teamQ) > 0 && len(taskQ) > 0 {
				tm0 = teamQ[0]
				tk0 = taskQ[0]
			}
			
			select {
			case tm := <-p.wholesaler.teams:
				teamQ = append(teamQ, tm)
			case tk := <-p.tasks:
				taskQ = append(taskQ, tk)
			case tm0.broker <- tk0:
				teamQ = teamQ[1:]
				taskQ = taskQ[1:]
			}
		}
	}()
}


func (p *parser) Parse(in interface{}) (interface{}, error) {
	p.tasks <- &task{id: 1, job: in}

	product := <- p.products
	if product.taskId != 1 {
		return nil, nil
	}
	return product.output, product.err
}