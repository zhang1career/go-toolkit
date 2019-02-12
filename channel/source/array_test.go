package source_test

import (
	"github.com/zhang1career/lib/channel/source"
	"testing"
)

func TestArraySource(t *testing.T) {
	in1 := []int{3, 2, 6, 7, 4}
	var inter1 = make([]interface{}, len(in1))
	for k, v := range in1 {
		inter1[k] = interface{}(v)
	}
	src1  := source.ArraySource(inter1)
	for v := range src1 {
		t.Log(v)
	}
}
