package tree

const (
	NODE_TYPE_LEAF = iota
	NODE_TYPE_BRANCH
)


type Node struct {
	children  []*Node
	val       interface{}
	typ       int
}

func CreateNode() Node {
	return Node{}
}

func (n *Node) GetChild(i int) *Node {
	return n.children[i]
}

func (n *Node) GetChildren() []*Node {
	return n.children
}

func (n *Node) AppendChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) SetLeaf(val interface{}) {
	n.val = val
	n.typ = NODE_TYPE_LEAF
}

func (n *Node) SetBranch(val interface{}) {
	n.val = val
	n.typ = NODE_TYPE_BRANCH
}

func (n *Node) GetValue() interface{} {
	return n.val
}
func (n *Node) IsLeaf() bool {
	return n.typ == NODE_TYPE_LEAF
}
