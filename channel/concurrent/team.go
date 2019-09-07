package concurrent

import "github.com/zhang1career/lib/channel/ctrlbus"

func createTeam(id int, work Work) *team {
	return &team{
		id: id,
		broker: make(chan *task),
		worker: worker{Work: work},
	}
}

func (t *team) run(ctrlbus *ctrlbus.Ctrlbus, wholesaler chan *team, products chan<- *product) {
	go func() {
		for {
			wholesaler <- t
			select {
			case <-ctrlbus.GetDone():
				return
			case task := <-t.broker:
				products <- &product{taskId: task.id, output: t.worker.Do(task.job)}
			}
		}
	}()
}
