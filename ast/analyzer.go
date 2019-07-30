package ast

import "strings"

var ks []Keyword

func Prepare() {
	ks = append(
		ks,
		CreateKeyword(`SELECT(.*)FROM(.*)WHERE(.*)`, nil),
	)
}

func Analyze(data []byte) *Node {
	Prepare()

	ret := CreateNode()

	var ms [][]byte
	var i int
	for j, k := range ks {
		if ms = k.Match(data); ms != nil {
			i = j
			break
		}
	}
	if ms == nil {
		ret.SetValue(strings.TrimSpace(string(data)))
		ret.SetTypeLeaf()
		return &ret
	}

	ret.SetValue(ks[i].GetObject())
	ret.SetTypeBranch()
	for _, m := range ms {
		ret.AppendChild(Analyze(m))
	}
	return &ret
}