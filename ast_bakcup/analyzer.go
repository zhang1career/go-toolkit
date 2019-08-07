package ast

import (
	"github.com/zhang1career/lib/ast/keyword"
	"github.com/zhang1career/lib/ast/tree"
	"strings"
)

var ks []operator.Keyword

func Prepare() {
	ks = append(
		ks,
		operator.CreateKeyword(`SELECT(.*)FROM(.*)WHERE(.*)`, nil),
	)
}

func Analyze(data []byte) *tree.Node {
	Prepare()

	ret := tree.CreateNode()

	var ms [][]byte
	var i int
	for j, k := range ks {
		if ms = k.Match(data); ms != nil {
			i = j
			break
		}
	}
	if ms == nil {
		ret.SetLeaf(strings.TrimSpace(string(data)))
		return &ret
	}

	ret.SetBranch(ks[i].GetObject())
	for _, m := range ms {
		ret.AppendChild(Analyze(m))
	}
	return &ret
}