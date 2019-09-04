package snowflake

import (
	"fmt"
)

func New(machine int, round int) *Snowflake {
	snow := Snowflake{
		machine: machine,
		round:   round,
	}
	snow.Reset()
	return &snow
}

func (this *Snowflake) Reset() {
	this.group = GetGroupBitcode(GetTime(), this.round, this.machine)
	this.serial = 0
}

func (this *Snowflake) Do() (uint64, error) {
	this.serial++
	if this.serial > SerialMax {
		return 0, fmt.Errorf("id sold out, please try later")
	}
	return SnowMake(this.group, this.serial), nil
}

func (this *Snowflake) Undo(id uint64) (int64, int, int, int) {
	time, round, machine, serial := SnowMelt(id)
	return int64(time + TimeStart), int(round), int(machine), int(serial)
}

//// round
//func (this *Snowflake) incRound() uint64 {
//	var round uint64
//	if this.round.Get().(int) < 0 {
//		round = this.round.Set(0).(uint64)
//	} else {
//		round = this.round.Inc().(uint64)
//	}
//	return (round << offsetRound) & maskRound
//}