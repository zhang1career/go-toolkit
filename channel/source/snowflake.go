package source

import (
	"math/rand"
	"time"
)

/* snowflake bitmap
 * msb                                                               lsb
 * 0-000000000000000000000000000000000000000-0000000000000000-00000000
 * |                    |                            |            |
 * reserved             createTime(t)                machineId(m) serialNo(s)
 *                      |                            |            |
 *                      +------------------ all -----+------------+
*/
const (
	bitAll       = 63
	bitT         = 39  // 2^39 / (1000*3600*24*365) = 17.4 year/(1 millisecond/bit)
	bitM         = bitAll - bitT - bitS
	bitS         = 12
	
	maskT        = uint64((1<<bitT - 1) << (bitM + bitS))
	maskM        = uint64((1<<bitM - 1) << bitS)
	maskS        = uint64(1<<bitS - 1)
	
	bitNano2Mil  = 23  // 10'000'000'000 nanosecond >> 23 = 1 ms
)

type Worker struct {
	t uint64
	m uint64
	s uint64
}

func CreateWorker(machineId uint16) *Worker {
	return &Worker{
		m:  uint64(machineId),
	}
}

func (this *Worker) Run(done <-chan interface{}) <-chan uint64{
	out := make(chan uint64)
	go func() {
		defer close(out)
		
		for {
			snowflake := this.Compose()
			select {
			case <-done:
				return
			case out <- snowflake:
			}
		}
	}()
	return out
}

func (this *Worker) Compose() uint64 {
	t := (uint64(time.Now().UnixNano()) << (bitM + bitS - bitNano2Mil)) & maskT
	m := (this.m << bitS) & maskM
	s := rand.Uint64() & maskS

	return uint64(t | m | s)
}

func (this *Worker) Decompose(snowflake uint64) Worker {
	return Worker{
		t: ((snowflake & maskT) >> (bitM + bitS - bitNano2Mil)) / 1e6, // in ms
		m: (snowflake & maskM) >> bitM,
		s: snowflake & maskS,
	}
}