package snowflake_test


//type Parser struct {
//	Output      chan interface{}
//	manager     []*ctrlbus.Ctrlbus
//	seekers     chan chan interface{}
//}
//
//func CreateParser() *Parser {
//	return &Parser{
//		Output:  make(chan interface{}),
//		manager: make([]*ctrlbus.Ctrlbus, 0),
//		seekers: make(chan chan interface{}),
//	}
//}
//
//func (this *Parser) AddWorker(num int, out chan int, back chan concurrent.Work) {
//	if num <= 0 {
//		return
//	}
//
//	for i := 0; i < num; i++ {
//		out <- i
//		work := <-back
//		worker := createWorker(i, work)
//		worker.run(this.seekers, this.Output)
//	}
//}
//
//func (this *Parser) Run(jobs chan interface{}) {
//	go func() {
//		var seekerQ = make([]chan interface{}, 0)
//		var jobQ = make([]interface{}, 0)
//
//		for {
//			var seeker chan interface{}
//			var job interface{}
//
//			if len(seekerQ) > 0 && len(jobQ) > 0 {
//				seeker = seekerQ[0]
//				job = jobQ[0]
//			}
//
//			select {
//			case s := <-this.seekers:
//				seekerQ = append(seekerQ, s)
//			case j := <-jobs:
//				jobQ = append(jobQ, j)
//			case seeker <- job:
//				seekerQ = seekerQ[1:]
//				jobQ = jobQ[1:]
//			}
//		}
//	}()
//}
//
//type worker struct {
//	id      int
//	concurrent.Work
//}
//
//func createWorker(id int, work concurrent.Work) *worker {
//	return &worker{id: id, Work: work}
//}
//
//func (this *worker) run(seekers chan chan interface{}, output chan<- interface{}) {
//	go func() {
//		var seeker = make(chan interface{})
//		defer close(seekers)
//		for {
//			seekers <- seeker
//			select {
//			case job := <-seeker:
//				out := this.Do(job)
//				output <- out
//			}
//		}
//	}()
//}
//
//
//func TestSnowFlakeQueue(t *testing.T) {
//	machineIds := make(chan int)
//	snows := snowflake.CreateSnowFlakeQueue(machineIds)
//	salt := make(chan interface{})
//
//	p := CreateParser()
//	p.AddWorker(3, machineIds, snows)
//	p.Run(salt)
//
//	salt <- uint64(0)
//	for i:=0; i<100; i++ {
//		v, ok := <-p.Output
//		if !ok {
//			t.Log("done")
//			return
//		}
//
//		time, machine, serial := snowflake.SnowMelt(v.(uint64))
//		t.Logf("%v, %03x, %03x", time, machine, serial)
//
//		salt <- v
//	}
//}