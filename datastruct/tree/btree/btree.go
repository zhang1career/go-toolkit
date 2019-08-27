package btree

import (
	"fmt"
)

type Node struct {
	value       int
	left, right *Node
}

func CreateNode(value int) (node *Node) {
	return &Node{value: value}
}

func (node *Node) Value() (value int) {
	if node == nil {
		return 0
	}
	return node.value
}

func (node *Node) SetValue(value int) (err error){
	if node == nil {
		return fmt.Errorf("node nil, SetValue failed")
	}
	node.value = value
	return nil
}
