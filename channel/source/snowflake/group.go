package snowflake

import (
	"fmt"
	"github.com/zhang1career/golab/cache"
	"github.com/zhang1career/golab/channel/concurrent"
	"github.com/zhang1career/golab/channel/ctrlbus"
	"github.com/zhang1career/golab/log"
	"github.com/zhang1career/golab/math/calc"
	"sync/atomic"
)

func CreateGroup(ctrlbus *ctrlbus.Ctrlbus, machine int) *SnowGroup {
	c := cache.NewOnce("localhost", 6379, nil)

	key := fmt.Sprintf("snow:m[%03d]:r", machine)
	round, ok := c.Incr(key); if !ok {
		log.Fatal("fail to get round")
	}

	s := SnowGroup{
		machine: machine,
		round:   int(round-1),
	}
	s.Reset()

	s.run(ctrlbus)

	return &s
}

func (s *SnowGroup) Reset() {
	s.group = GetGroupBitcode(s.round, GetTime(), s.machine)
	s.serial = 0
}

func (s *SnowGroup) run(ctrlbus *ctrlbus.Ctrlbus) {
	go func() {
		for {
			select {
			case <- ctrlbus.GetDone():
				return
			case <- ctrlbus.GetTicker().C:
				s.Reset()
			}
		}
	}()
}

func CreateGroupAsWorker(ctrlbus *ctrlbus.Ctrlbus, machine int) concurrent.Work {
	return CreateGroup(ctrlbus, machine)
}

func (s *SnowGroup) Do(n interface{}) concurrent.Output {
	tmp := atomic.LoadInt32(&s.serial)
	if tmp < 0 || tmp > SerialMax {
		return concurrent.CreateOutput(nil, fmt.Errorf("id sold out, please try later"))
	}

	m := calc.MinInt32(int32(n.(int)), SerialMax-tmp+1)
	if !atomic.CompareAndSwapInt32(&s.serial, tmp, tmp+m) {
		return concurrent.CreateOutput(nil, fmt.Errorf("fail to increase serial[%d], please try later", tmp))
	}

	ret := make([]uint64, m)
	for i := 0; i < int(m); i++ {
		ret[i] = SnowMake(s.group, int(tmp)+i)
	}
	return concurrent.CreateOutput(ret, nil)
}