package concurrent

import "github.com/zhang1career/lib/channel/ctrlbus"

func createWorker(id int, work Work) *worker {
	return &worker{
		id: id,
		Work: work,
	}
}

func (w *worker) run(ctrlbus *ctrlbus.Ctrlbus, seekers chan chan interface{}, output chan<- interface{}) {
	go func() {
		var seeker = make(chan interface{})
		defer close(seeker)
		for {
			seekers <- seeker
			select {
			case <-ctrlbus.GetDone():
				return
			case job := <-seeker:
				output <- w.Do(job)
			}
		}
	}()
}
