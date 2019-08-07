package ast_test

import (
	"testing"
)

func TestNewKeyword(t *testing.T) {
	k := CreateKeyword(`SELECT(.*)FROM(.*)WHERE(.*)`, nil)
	t.Log(k)
}

func TestKeyword_Match(t *testing.T) {
	k := CreateKeyword(`SELECT(.*)FROM(.*)WHERE(.*)`, nil)
	sql := []byte("SELECT * FROM rules WHERE Id=2")
	ret := k.Match(sql)

	field := string(ret[0][:])
	t.Log(field)
	table := string(ret[1][:])
	t.Log(table)
	cond := string(ret[2][:])
	t.Log(cond)
}

func TestKeyword_NoMatch(t *testing.T) {
	k := CreateKeyword(`SELECT(.*)FROM(.*)WHERE(.*)`, nil)
	sql := []byte("INSERT IN rules Id=2")
	ret := k.Match(sql)

	t.Log(ret)
}
