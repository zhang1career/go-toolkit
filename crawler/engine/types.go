package engine

type Request struct {
	Url   string
	Parse func([]byte) ParseResult
}

type ExploreResult struct {
	Data  []byte
	Parse func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}
