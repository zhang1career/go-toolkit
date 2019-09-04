package snowflake

import (
	"github.com/zhang1career/lib/log"
	"time"
)

/*
 * base function
 */
func SnowMake(group uint64, serial int) uint64 {
	return group | (uint64(serial) & maskSerial)
}

func GetGroupBitcode(time int64, round int, machine int) uint64 {
	return (uint64(time)    << offsetTime) & maskTime  |
		   (uint64(round)   << offsetRound) & maskRound |
		   (uint64(machine) << offsetMachine) & maskMachine
}

func GetTime() int64 {
	now := time.Now().Unix()
	if now < TimeStart {
		tm := time.Unix(TimeStart, 0)
		log.Fatal("System time should not be early than %s", tm.Format("2006-01-02 03:04:05 PM"))
	}
	return now - TimeStart
}

func SnowMelt(input uint64) (uint64, uint64, uint64, uint64) {
	t := (input & maskTime) >> offsetTime
	r := (input & maskRound) >> offsetRound
	m := (input & maskMachine) >> offsetMachine
	s := input & maskSerial
	return t, r, m, s
}