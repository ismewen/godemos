package search

type Node struct {
	value int
	next  *Node
}

func (node *Node) Insert(n *Node) {
	if node.next != nil {
		n.next, node.next = node.next, n
	} else {
		node.next = n
	}
}

func (node *Node) Get(value int) *Node {
	n := node
	for n.next != nil {
		if n.value == value {
			return n
		}
		n = n.next
	}
	return nil
}
