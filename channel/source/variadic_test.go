package source_test

import (
	"github.com/zhang1career/lib/channel/source"
	"testing"
)

func TestVariadicSource(t *testing.T) {
	src1  := source.VariadicSource(3,2,6,7,4)
	for v := range src1 {
		t.Log(v)
	}
	
	src2  := source.VariadicSource(9,1,8,0,5)
	for v := range src2 {
		t.Log(v)
	}
}
