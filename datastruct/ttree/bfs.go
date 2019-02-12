package ttree

type BFS struct {
	Traverser
}

func NewBFS(
	explore func(interface{}) (interface{}, error),
	parse   func(interface{}) ([]interface{}, interface{}),
	root    func(interface{}) bool,
	done    func(interface{}) bool) *BFS {
	b := &BFS{
		Traverser: Traverser{
			Explore: explore,
			Parse: parse,
			Root: root,
			Done: done,
		},
	}
	b.pend = b.pending
	return b
}

func (s *BFS) pending(slice[]interface{}, elems interface{}) []interface{} {
	return append(slice, elems)
}
