package source_test

import (
	"github.com/zhang1career/golab/channel/source"
	"testing"
)

func TestRandSource(t *testing.T) {
	src1  := source.RandSource(5)
	for v := range src1 {
		t.Log(v)
	}
	
	src2  := source.RandSource(5)
	for v := range src2 {
		t.Log(v)
	}
	
}