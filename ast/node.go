package ast

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

func (n *Node) SetChild(i int, child *Node) {
	n.children[i] = child
}

func (n *Node) AppendChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) GetValue() interface{} {
	return n.val
}

func (n *Node) SetValue(val interface{}) {
	n.val = val
}

func (n *Node) IsTypeLeaf() bool {
	return n.typ == NODE_TYPE_LEAF
}

func (n *Node) SetTypeLeaf() {
	n.typ = NODE_TYPE_LEAF
}

func (n *Node) SetTypeBranch() {
	n.typ = NODE_TYPE_BRANCH
}