package operation_test

import (
	"github.com/zhang1career/lib/ast/operation"
	"testing"
)

var sel = map[string]interface{} {
	"select": map[string]interface{} {
		"target": "rules",
		"source": "",
		"cond":   "id=2",
	},
}

func TestNew(t *testing.T) {
	op := operation.New(sel)
	t.Log(op)
}
