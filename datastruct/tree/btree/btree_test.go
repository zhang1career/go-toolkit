package btree

import "fmt"

func print(node *Node) {
	if node == nil {
		fmt.Print(nil)
		return
	}
	fmt.Println(node.value)
}

func ExampleNode_TraverseWithFunc() {
	root := CreateNode(1)
	root.left = CreateNode(2)
	root.right = CreateNode(3)
	root.TraverseWithFunc(print)
	
	
	// output:
	// 2
	// 1
	// 3
}

func ExampleNode_TraverseByChannel() {
	root := CreateNode(11)
	root.left = CreateNode(222)
	root.right = CreateNode(3)
	
	c := root.TraverseToChannel()
	maxNode := 0
	for node := range c {
		if node.Value() > maxNode {
			maxNode = node.Value()
		}
	}
	fmt.Println(maxNode)
	
	// output:
	// 222
}

