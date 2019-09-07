package concurrent

import (
	"github.com/zhang1career/lib/channel/ctrlbus"
)

func CreateParser() *parser {
	return &parser{
		managers:   make([]*ctrlbus.Ctrlbus, 0),
		wholesaler: make(chan *team),
		products:   make(chan *product),
	}
}

func (p *parser) AddTeam(config map[string]interface{}, num int, f func(i int) Work) {
	if num <= 0 {
		return
	}

	for i := len(p.managers); i < len(p.managers) + num; i++ {
		manager := ctrlbus.CreateCtrlbus(config)
		p.managers = append(p.managers, manager)

		team := createTeam(i, f(i))
		team.run(manager, p.wholesaler, p.products)
	}
}

func (p *parser) Run(tasks chan *task) {
	go func() {
		var teamQ = make([]*team, 0)
		var taskQ = make([]*task, 0)
		
		for {
			var team *team
			var task *task
			
			if len(teamQ) > 0 && len(taskQ) > 0 {
				team = teamQ[0]
				task = taskQ[0]
			}
			
			select {
			case tm := <-p.wholesaler:
				teamQ = append(teamQ, tm)
			case tk := <-tasks:
				taskQ = append(taskQ, tk)
			case team.broker <- task:
				teamQ = teamQ[1:]
				taskQ = taskQ[1:]
			}
		}
	}()
}
