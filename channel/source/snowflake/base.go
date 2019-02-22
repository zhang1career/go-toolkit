package snowflake

import (
	"math/rand"
	"time"
)

/*
 * base function
 */
func SnowMake(machineId uint64) uint64 {
	t := (uint64(time.Now().UnixNano()) << (bitM + bitS - bitNano2Mil)) & maskT
	m := (machineId << bitS) & maskM
	s := uint64(rand.Uint32()) & maskS
	
	return t | m | s
}

func SnowMelt(input uint64) (int, int, int) {
	t := int(((input & maskT) >> (bitM + bitS - bitNano2Mil)) / 1e6) // in ms
	m := int((input & maskM) >> bitM)
	s := int(input & maskS)
	return t, m, s
}