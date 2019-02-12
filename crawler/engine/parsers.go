package engine

// default parser
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
