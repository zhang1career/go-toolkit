package snowflake

import (
	"fmt"
	"github.com/zhang1career/lib/cache"
	"github.com/zhang1career/lib/log"
	"sync/atomic"
)

func New(machine int) *Snowflake {
	c := cache.NewOnce("localhost", 6379, nil)

	key := fmt.Sprintf("snow:m[%0d]:r", machine)
	round, ok := c.Incr(key); if !ok {
		log.Fatal("fail to get round")
	}

	snow := Snowflake{
		machine: machine,
		round:   int(round-1),
	}
	snow.Reset()
	return &snow
}

func (this *Snowflake) Reset() {
	this.group = GetGroupBitcode(this.round, GetTime(), this.machine)
	this.serial = 0
}

func (this *Snowflake) Do() (uint64, error) {
	tmp := atomic.LoadInt32(&this.serial)
	if tmp < 0 || tmp > SerialMax {
		return 0, fmt.Errorf("id sold out, please try later")
	}
	if !atomic.CompareAndSwapInt32(&this.serial, tmp, tmp+1) {
		return 0, fmt.Errorf("fail to increase serial, serial[%d]", tmp)
	}
	return SnowMake(this.group, int(tmp)), nil
}