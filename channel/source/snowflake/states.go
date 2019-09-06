package snowflake

import (
	"github.com/zhang1career/lib/log"
	"time"
)

func (this *SnowGroup) Run(done <-chan interface{}, tic *time.Ticker, req <-chan int32) <-chan []uint64 {
	out := make(chan []uint64)
	go func() {
		defer close(out)
		for {
			select {
				case <- done:
					return
				case <- tic.C:
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
