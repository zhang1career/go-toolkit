package ttree

import "github.com/zhang1career/lib/datastruct/dim/unidim"

type DFS struct {
	Traverser
}

func NewDFS(
	explore func(interface{}) (interface{}, error),
	parse   func(interface{}) ([]interface{}, interface{}),
	root    func(interface{}) bool,
	done    func(interface{}) bool) *DFS {
	d := &DFS{
		Traverser: Traverser{
			Explore: explore,
			Parse: parse,
			Root: root,
			Done: done,
		},
	}
	d.pend = d.pendding
	return d
}

func (s *DFS) pendding(slice[]interface{}, elems interface{}) []interface{} {
	return unidim.Prepend(slice, elems)
}