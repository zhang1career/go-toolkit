package snowflake

type SnowGroup struct {
	machine int
	round   int
	serial  int32
	group   uint64
}
