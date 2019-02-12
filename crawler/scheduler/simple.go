package scheduler

type SinmpeScheduler struct {
	requests chan interface{}
}

func (s *SinmpeScheduler) Config(conf map[string]interface{}) {
	s.requests = (conf["request_chan"]).(chan interface{})
}

func (s *SinmpeScheduler) Submit(r interface{}) {
	go func() {
		s.requests <- r
	}()
}
