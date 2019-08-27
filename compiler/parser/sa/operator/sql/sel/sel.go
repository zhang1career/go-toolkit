package sel

import "github.com/zhang1career/lib/compiler/parser/sa"

type Sel struct {
	value       string
	preCount    int
	postCount   int
}

func New() sa.Operator {
	return &Sel{
		value:      "select",
		preCount:   0,
		postCount:  1,
	}
}

func (this *Sel) GetValue() string {
	return this.value
}

func (this *Sel) GetPreCount() int {
	return this.preCount
}

func (this *Sel) GetPostCount() int {
	return this.postCount
}
