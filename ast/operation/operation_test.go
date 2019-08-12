package operation_test

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/ast/operation"
	"testing"
)

func TestSelect(t *testing.T) {
	var data = ast.Item {
		"select": ast.Item {
			"target": "rules",
			"source": "risk",
			"cond":   "id=2",
		},
	}
	
	op   := operation.New(data)
	test := op.Evaluate()
	comp := "/risk/rules?id=2"
	if test != comp {
		t.Errorf("Evaluation was incorrect, got: %s, want: %s.", test, comp)
	}
}