package operation_test

import (
	"github.com/zhang1career/lib/compiler"
	"github.com/zhang1career/lib/compiler/ast/operation"
	"testing"
)

func TestAdd(t *testing.T) {
	var data = compiler.Dim {
		"+": []interface{} {
			1,
			2,
			3,
		},
	}
	
	op   := operation.ImportFromMap(data)
	got  := op.Evaluate()
	want := 6
	if got != want {
		t.Errorf("Evaluation was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestGreatThan(t *testing.T) {
	var data = compiler.Dim {
		">": []interface{} {
			2,
			1,
		},
	}

	op   := operation.ImportFromMap(data)
	got  := op.Evaluate()
	want := true
	if got != want {
		t.Errorf("Evaluation was incorrect, got: %t, want: %t.", got, want)
	}
}

func TestSelect(t *testing.T) {
	var data = compiler.Dim {
		"select": compiler.Dim {
			"target":   "rules",
			"from":     "risk",
			"where":    "id=2",
		},
	}

	op   := operation.ImportFromMap(data)
	got  := op.Evaluate()
	want := "/risk/rules?id=2"
	if got != want {
		t.Errorf("Evaluation was incorrect, got: %s, want: %s.", got, want)
	}
}