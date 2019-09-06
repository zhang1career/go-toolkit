package snowflake

import (
	"github.com/zhang1career/lib/channel/concurrent"
	"time"
)

func CreateQueue(machines chan int) chan concurrent.Work {
	output := make(chan concurrent.Work)
	go func() {
		defer close(output)
		for {
			select {
			case m := <-machines:
				output <- CreateGroup(m)
			}
		}
	}()
	return output





	done := make(chan interface{})
	defer close(done)

	tic  := time.NewTicker(time.Second)

	req  := make(chan int32)
	defer close(req)

	groups := make([]*SnowGroup, 0)
	ret := make();
	for _, id := range machines {
		if id < 0 || id > MachineMax {
			continue
		}
		g := CreateGroup(id)
		got := g.Run(done, tic, req)

		time.Sleep(time.Second)
		req <- 1
		t.Logf("%0x\r", <-got)
	}
}