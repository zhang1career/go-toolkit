package engine

import (
	"fmt"
	"github.com/zhang1career/lib/datastruct/tree/ttree"
	"github.com/zhang1career/lib/log"
)

type SimpleEngine struct {
}

func (e *SimpleEngine) start(seed ...Request) []interface{} {
	roots := make([]interface{}, 0)
	for _, s := range seed {
		roots = append(roots, s)
	}
	return roots
}

func (e *SimpleEngine) Run(seed ...Request) {
	w := worker{dumps:make(map[string]bool)}
	b := ttree.NewBFS(w.explore, w.parse, w.root, w.done)
	
	var out []interface{}
	_, err := b.Traverse(e.start(seed...), &out)
	if err != nil {
		log.Error(err.Error())
		return
	}
	
	for o := range out {
		fmt.Println(o)
	}
}
