package operation_test

import (
	"github.com/zhang1career/lib/ast"
	"github.com/zhang1career/lib/ast/operation"
	"testing"
)

func TestAdd(t *testing.T) {
	var data = ast.Item {
		"+": []interface{} {
			1,
			2,
			3,
		},
	}
	
	op   := operation.New(data)
	got  := op.Evaluate()
	want := 6
	if got != want {
		t.Errorf("Evaluation was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestGreatThen(t *testing.T) {
	var data = ast.Item {
		">": []interface{} {
			2,
			1,
		},
	}

	op   := operation.New(data)
	got  := op.Evaluate()
	want := true
	if got != want {
		t.Errorf("Evaluation was incorrect, got: %t, want: %t.", got, want)
	}
}

func TestSelect(t *testing.T) {
	var data = ast.Item {
		"select": ast.Item {
			"target":   "rules",
			"from":     "risk",
			"where":    "id=2",
		},
	}

	op   := operation.New(data)
	got  := op.Evaluate()
	want := "/risk/rules?id=2"
	if got != want {
		t.Errorf("Evaluation was incorrect, got: %s, want: %s.", got, want)
	}
}