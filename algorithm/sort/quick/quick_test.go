package quick_test

import (
	"github.com/zhang1career/golab/algorithm/sort/quick"
	"testing"
)

func TestRun(t *testing.T) {
	params := []int{-4,0,7,4,9,-5,-1,0,-7,-1}
	quick.Run(params)
	t.Log(params)
}