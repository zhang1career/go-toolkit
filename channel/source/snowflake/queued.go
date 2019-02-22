package snowflake

import (
	"github.com/zhang1career/lib/channel/concurrent"
)

/*
 * queue function
 */
func CreateSnowFlakeQueue(machineIds chan int) chan concurrent.Work {
	output := make(chan concurrent.Work)
	go func() {
		defer close(output)
		for {
			select {
			case id := <-machineIds:
				output <- CreateSnowFlake(id)
			}
		}
	}()
	return output
}