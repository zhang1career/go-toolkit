package snowflake

import (
	"math/rand"
)

/*
 * stateful function
 */
type Snowflake struct {
	m   interface{}
}

func CreateSnowFlake(machineId int) *Snowflake {
	return &Snowflake{m: uint64(machineId)}
}

func (this *Snowflake) Do(seed interface{}) interface{} {
	rand.Seed(int64(seed.(uint64)))
	return SnowMake(this.m.(uint64))
}

func (this *Snowflake) Undo(input interface{}) (int, int, int) {
	return SnowMelt(input.(uint64))
}