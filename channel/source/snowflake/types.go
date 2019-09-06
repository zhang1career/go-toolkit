package snowflake

type Snowflake struct {
	machine int
	round   int
	serial  int32
	group   uint64
}
