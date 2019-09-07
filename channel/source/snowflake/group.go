package snowflake

import (
	"fmt"
	"github.com/zhang1career/lib/cache"
	"github.com/zhang1career/lib/channel/concurrent"
	"github.com/zhang1career/lib/log"
	"github.com/zhang1career/lib/math/calc"
	"sync/atomic"
)

func CreateGroup(machine int) concurrent.Work {
	c := cache.NewOnce("localhost", 6379, nil)

	key := fmt.Sprintf("snow:m[%0d]:r", machine)
	round, ok := c.Incr(key); if !ok {
		log.Fatal("fail to get round")
	}

	sg := SnowGroup{
		machine: machine,
		round:   int(round-1),
	}
	sg.Reset()
	return &sg
}

func (this *SnowGroup) Reset() {
	this.group = GetGroupBitcode(this.round, GetTime(), this.machine)
	this.serial = 0
}

func (this *SnowGroup) Do(n interface{}) (interface{}, error) {
	tmp := atomic.LoadInt32(&this.serial)
	if tmp < 0 || tmp > SerialMax {
		return nil, fmt.Errorf("id sold out, please try later")
	}

	m := calc.MinInt32(n, SerialMax-tmp+1)
	if !atomic.CompareAndSwapInt32(&this.serial, tmp, tmp+m) {
		return nil, fmt.Errorf("fail to increase serial[%d], please try later", tmp)
	}

	ret := make([]uint64, m)
	for i := 0; i < int(m); i++ {
		ret[i] = SnowMake(this.group, int(tmp)+i)
	}
	return ret, nil
}