package concurrent_test

import (
	"github.com/zhang1career/lib/channel/concurrent"
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
)

func TestParser(t *testing.T) {
	machineIds := make(chan int)
	snows := snowflake.CreateSnowFlakeQueue(machineIds)
	salt := make(chan interface{})
	
	p := concurrent.CreateParser()
	p.AddWorker(3, snows, machineIds)
	p.Run(salt)
	
	salt <- uint64(0)
	for i:=0; i<100; i++ {
		v, ok := <-p.Output
		if !ok {
			t.Log("done")
			return
		}
		
		time, machine, serial := snowflake.SnowMelt(v.(uint64))
		t.Logf("%v, %03x, %03x", time, machine, serial)
		
		salt <- v
	}
}
