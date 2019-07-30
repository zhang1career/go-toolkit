package ast_test

import (
	"github.com/zhang1career/lib/ast"
	"testing"
)

func TestNewKeyword(t *testing.T) {
	k := ast.CreateKeyword(`SELECT(.*)FROM(.*)WHRER(.*)`, nil)
	t.Log(k)
}

func TestKeyword_Match(t *testing.T) {
	k := ast.CreateKeyword(`SELECT(.*)FROM(.*)WHRER(.*)`, nil)
	sql := []byte("SELECT * FROM rules WHRER Id=2")
	ret := k.Match(sql)

	field := string(ret[0][:])
	t.Log(field)
	table := string(ret[1][:])
	t.Log(table)
	cond := string(ret[2][:])
	t.Log(cond)
}

func TestKeyword_NoMatch(t *testing.T) {
	k := ast.CreateKeyword(`SELECT(.*)FROM(.*)WHRER(.*)`, nil)
	sql := []byte("INSERT IN rules Id=2")
	ret := k.Match(sql)

	t.Log(ret)
}
