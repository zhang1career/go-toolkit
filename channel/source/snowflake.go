package source

import (
	"github.com/zhang1career/lib/channel/concurrent"
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


/*
 * base function
 */
func Compose(machineId uint64) uint64 {
	t := (uint64(time.Now().UnixNano()) << (bitM + bitS - bitNano2Mil)) & maskT
	m := (machineId << bitS) & maskM
	s := uint64(rand.Uint32()) & maskS
	
	return t | m | s
}

func Decompose(input uint64) (int, int, int) {
	t := int(((input & maskT) >> (bitM + bitS - bitNano2Mil)) / 1e6) // in ms
	m := int((input & maskM) >> bitM)
	s := int(input & maskS)
	return t, m, s
}


/*
 * stateful function
 */
type snowflake struct {
	m   interface{}
}

func CreateSnowFlake(machineId int) *snowflake {
	return &snowflake{m: uint64(machineId)}
}

func (this *snowflake) Do(seed interface{}) interface{} {
	rand.Seed(int64(seed.(uint64)))
	return Compose(this.m.(uint64))
}

func (this *snowflake) Undo(input interface{}) (int, int, int) {
	return Decompose(input.(uint64))
}


/*
 * flow function
 */
func CreateSnowFlow(machineIds chan int) chan concurrent.Work {
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