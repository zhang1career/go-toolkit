package popup_test

import (
	"github.com/zhang1career/golab/algorithm/sort/popup"
	"testing"
)

func TestRun(t *testing.T) {
	params := []int{3, 2, 1}
	popup.Run(&params)
	t.Log(params)
}