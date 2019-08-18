package mtree

import (
	"fmt"
	"github.com/zhang1career/lib/gotime"
)

type Node struct {
	value       interface{}
	parent      *Node
	children    []*Node
}

func New(value interface{}) *Node {
	return &Node{value: value}
}

func (this *Node) SetValue(value interface{}) error {
	if err := this.check(); err != nil {
		return err
	}
	this.value = value
	return nil
}

func (this *Node) GetValue() (interface{}, error) {
	if err := this.check(); err != nil {
		return nil, err
	}
	return this.value, nil
}

func (this *Node) SetParent(parent *Node) error {
	if err := this.check(); err != nil {
		return err
	}
	this.parent = parent
	return nil
}

func (this *Node) GetParent() (*Node, error) {
	if err := this.check(); err != nil {
		return nil, err
	}
	return this.parent, nil
}

func (this *Node) SetChild(child *Node) error {
	if err := this.check(); err != nil {
		return err
	}
	this.children = append(this.children, child)
	return nil
}

func (this *Node) GetChild(index int) (*Node, error) {
	if err := this.check(); err != nil {
		return nil, err
	}
	if index < 0 || index >= len(this.children) {
		return nil, fmt.Errorf("Index is out of range, %s failed\n", gotime.ThisFunc())
	}
	return this.children[index], nil
}

func (this *Node) check() error {
	if this == nil {
		return fmt.Errorf("Tree node is nil, %s failed\n", gotime.ThisFunc())
	}
	return nil
}

func (this *Node) GetGrandestParent() (*Node, error) {
	if err := this.check(); err != nil {
		return nil, err
	}
	if this.parent == nil {
		return this, nil
	}
	return this.parent.GetGrandestParent()
}

func (this *Node) AddParent(parent *Node) (*Node, error) {
	if err := this.check(); err != nil {
		return nil, err
	}
	grand, err := this.GetGrandestParent()
	if err != nil {
		return this, err
	}
	err = grand.SetParent(parent)
	if err != nil {
		return this, err
	}
	return parent, nil
}