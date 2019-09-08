package snowflake

import "github.com/zhang1career/lib/channel/concurrent"

type SnowGroup struct {
	machine int
	round   int
	serial  int32
	group   uint64
}

type SnowQueue struct {
	parser  *concurrent.Parser
}