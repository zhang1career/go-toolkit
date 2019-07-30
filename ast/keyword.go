package ast

import (
	"regexp"
)

type Keyword struct {
	reg   *regexp.Regexp
	obj   *KeyObject
}

type KeyObject interface {
	Evaluate() []byte
}

func CreateKeyword(pattern string, obj *KeyObject) Keyword {
	return Keyword{
		reg: regexp.MustCompile(pattern),
		obj: obj,
	}
}

func (k *Keyword) Match(data []byte) [][]byte {
	ms := k.reg.FindSubmatch(data)
	if ms == nil {
		return nil
	}
	return ms[1:]
}

func (k *Keyword) GetObject() *KeyObject {
	return k.obj
}