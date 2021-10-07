package tree

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Name  string
}

// 六种遍历方法 遍历方式:前序，中序，后序, 实现方式: 递归和循环实现

func NewTestTree() *Node {

	head := Node{Name: "A"}
	left := Node{Name: "B"}
	right := Node{Name: "C"}
	head.Left = &left
	head.Right = &right

	leftLeft := Node{Name: "D"}
	leftRight := Node{Name: "E"}

	left.Left = &leftLeft
	left.Right = &leftRight

	rightLeft := Node{Name: "F"}
	rightRight := Node{Name: "G"}

	right.Left = &rightLeft
	right.Right = &rightRight
	return &head

}

func RecursionPreOrder(node *Node) {
	// 前序遍历， 根 -> 左 ->右
	fmt.Println(node.Name)
	RecursionPreOrder(node.Left)
	RecursionPreOrder(node.Right)
}

func RecursionPostOrder(node *Node) {
	// hz
	RecursionPostOrder(node.Left)
	fmt.Print()
}
