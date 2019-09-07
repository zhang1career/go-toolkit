package snowflake

import (
	"github.com/zhang1career/lib/channel/ctrlbus"
	"github.com/zhang1career/lib/log"
)

func (this *SnowGroup) Run(ctrlbus *ctrlbus.Ctrlbus, req <-chan int32) <-chan []uint64 {
	out := make(chan []uint64)
	go func() {
		defer close(out)
		for {
			select {
				case <- ctrlbus.GetDone():
					return
				case <- ctrlbus.GetTicker().C:
					this.Reset()
				case n := <- req:
					ids, err := this.Do(n); if err != nil {
						log.Error(err.Error())
						break
					}
					out <- ids
			}
		}
	}()
	return out
}
