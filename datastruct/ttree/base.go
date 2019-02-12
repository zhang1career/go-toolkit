package ttree

// tree traverser
//   Explore: explore a node, return result and error
//   Parse:   parse result, turn into child nodes and valuable info
//   Root:    mark a node as root, and return whether the node had been root before
//   Done:    return whether traversal terminates or continues
type Traverser struct {
	Explore func(interface{}) (interface{}, error)
	Parse   func(interface{}) ([]interface{}, interface{})
	Root    func(interface{}) bool
	Done    func(interface{}) bool
	pend    func([]interface{}, interface{}) []interface{}
}

// Tree traversal with Band First Searc.
// param
//   input:
//   output:
func (s *Traverser) Traverse(input []interface{}, output *[]interface{}) (bool, error) {
	// check param
	searchingNodes := s.check(&input)
	// init searching pool
	var node interface{}
	
	// traverse
	for len(searchingNodes) > 0 {
		// pop a searching root
		node, searchingNodes = searchingNodes[0], searchingNodes[1:]
		// mark root
		s.Root(node)
		// exit due to condition
		if s.Done(node) {
			return true, nil
		}
		// explore the tree
		result, err := s.Explore(node)
		if err != nil {
			continue
		}
		// parse result
		children, out := s.Parse(result)
		// push to searching pool
		for _, child := range children {
			if child == nil || s.Root(child) {
				continue
			}
			searchingNodes = s.pend(searchingNodes, child)
		}
		// output
		//@todo 这一句 append 好像有问题，for range output 打印出来不是过程数组，而是结果数组
		*output = append(*output, out)
	}
	// complete
	return false, nil
}

func (s *Traverser) check(seeds *[]interface{}) []interface{} {
	var nodes []interface{}
	
	for _, s := range *seeds {
		if s == nil {
			continue
		}
		nodes = append(nodes, s)
	}

	return nodes
}