package engine

func index(node interface{}) Request {
	return node.(Request)
}
