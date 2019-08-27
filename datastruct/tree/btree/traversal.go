package btree

func (node *Node) TraverseWithFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.left.TraverseWithFunc(f)
	f(node)
	node.right.TraverseWithFunc(f)
}

func (node *Node) TraverseToChannel() (out chan<- *Node) {
	out = make(chan *Node)
	go func() {
		node.TraverseWithFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
