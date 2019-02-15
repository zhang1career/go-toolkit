package concurrent

import "github.com/zhang1career/lib/channel/ctrlbus"

type Work interface {
	Do(interface{}) interface{}
}

type worker struct {
	id      int
	Work
}

func createWorker(id int, work Work) *worker {
	return &worker{id: id, Work: work}
}

func (this *worker) run(ctrlbus *ctrlbus.Ctrlbus, seekers chan chan interface{}, output chan<- interface{}) {
	go func() {
		var seeker = make(chan interface{})
		defer close(seekers)
		for {
			seekers <- seeker
			select {
			case <-ctrlbus.Done:
				return
			case job := <-seeker:
				output <- this.Do(job)
			}
		}
	}()
}
