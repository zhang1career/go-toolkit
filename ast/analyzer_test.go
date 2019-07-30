package ast_test

import (
	"github.com/zhang1career/lib/ast"
	"testing"
)

func TestAnalyze(t *testing.T) {
	sql := []byte("SELECT * FROM rules WHERE Id=2")
	ret := ast.Analyze(sql)
	t.Log(ret)
}