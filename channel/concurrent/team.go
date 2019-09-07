package concurrent

import "github.com/zhang1career/lib/channel/ctrlbus"

func createTeam(id int, work Work) *team {
	return &team{
		id: id,
		broker: make(chan *task),
		worker: &worker{Work: work},
	}
}

func (t *team) run(ctrlbus *ctrlbus.Ctrlbus, wholesaler *wholesaler, products chan<- *product) {
	go func() {
		for {
			wholesaler.teams <- t
			select {
			case <-ctrlbus.GetDone():
				return
			case task := <-t.broker:
				output, err := t.worker.Do(task.job)
				products <- &product{taskId: task.id, output: output, err: err}
			}
		}
	}()
}
