package stack

import (
	"github.com/zhang1career/lib/log"
)

type Stack []interface{}

func New() *Stack {
	return &Stack{}
}

func (this *Stack) Push(v interface{}) {
	if this == nil {
		log.Fatal("Stack is nil, Push failed")
	}
	*this = append(*this, v)
}

func (this *Stack) Pop() interface{} {
	if this == nil {
		log.Fatal("Stack is nil, Pop failed")
	}
	n := len(*this) - 1
	if n < 0 {
		log.Warn("Stack is empty, Pop failed")
		return nil
	}
	
	tail := (*this)[n]
	*this = (*this)[:n]
	return tail
}

func (this *Stack) IsEmpty() (isEmpty bool) {
	return len(*this) <= 0
}
