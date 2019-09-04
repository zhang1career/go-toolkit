package snowflake

type Snowflake struct {
	machine int
	round   int
	serial  int
	group   uint64
}
