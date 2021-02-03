package search

type BinarySearchNode struct {
	left  *BinarySearchNode
	right *BinarySearchNode
	N     int
	Key   string
	value int
}

func Size(node *BinarySearchNode) int {
	if node ==  nil {
		return 0
	}
	return Size(node.left) + Size(node.right) + 1
}

func Min(node *BinarySearchNode) *BinarySearchNode {
	if node.left == nil {
		return node
	}
	return Min(node.left)
}

func DeleteMin(node *BinarySearchNode)  *BinarySearchNode {
	if node.left == nil {
		return node.right
	}
	node.left = DeleteMin(node.left)
	node.N = Size(node.left)  + Size(node.right) + 1
	return node
}

func DeleteMax(node *BinarySearchNode) *BinarySearchNode {
	if node.right == nil {
		return node.left
	}
	node.right = DeleteMax(node.right)
	node.N = Size(node.left)  + Size(node.right) + 1
	return node
}

func Delete(node *BinarySearchNode, key string) *BinarySearchNode{
	if node == nil {
		return nil
	}
	if node.Key > key{
		// 在左边进行删除
		node.left = Delete(node.left, key)
	}else if node.Key < key{
		// 在左边进行删除
		node.right = Delete(node.right, key)
	}else{
		// 此时node.key == key
		if node.left == nil {
			// parent指向node.right, node 被孤立，被回收
			return node.right
		}else if node.right == nil {
			// parent 指向了node.left node被孤立，被回收
			return node.left
		}else{
			t := Min(node.right)
			t.right = DeleteMin(node.right)
			t.left = node.left
		}
	}
	node.N = Size(node.left) + Size(node.right) + 1
	return node
}