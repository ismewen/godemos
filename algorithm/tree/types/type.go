package types

import (
	qtypes "algorithm/queue/types"
	"algorithm/stack/types"
	"errors"
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Name  string
}

// 六种遍历方法 遍历方式:前序，中序，后序, 实现方式: 递归和循环实现
// 前序，中序， 后序根据跟节点划分, 左优先于右输出
// 层序遍历
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

func RecursionPreOrder(node *Node, res *[]string) {
	// 递归前序遍历， 根 -> 左 ->右
	if node == nil {
		return
	}
	fmt.Print(node.Name)
	*res = append(*res, node.Name)
	RecursionPreOrder(node.Left, res)
	RecursionPreOrder(node.Right, res)
}

func RecursionPostOrder(node *Node) {
	// 递归后序遍历， 左 -> 右 -> 根
	if node == nil {
		return
	}
	RecursionPostOrder(node.Left)
	RecursionPostOrder(node.Right)
	fmt.Print(node.Name)
}

func RecursionInOrder(node *Node) {
	// 递归中序遍历, 左 > 根 -> 右
	if node == nil {
		return
	}
	RecursionInOrder(node.Left)
	fmt.Print(node.Name)
	RecursionInOrder(node.Right)
}

func LoopPreOrder(node *Node) {
	// 1.给定一个节点p，访问并入栈, p的左节点赋值给p, 不断重复直到p为nil
	// 判断 栈是否为空，弹出 栈顶 元素, 赋值给p
	// 不断重复，直到p和栈都为空
	stack := types.NewStack()
	p := node
	for p != nil || !stack.Empty() {
		// 左节点依次入栈
		for p != nil {
			// 访问p, 左节点一次入栈
			fmt.Print(p.Name)
			stack.Push(p)
			p = p.Left
		}
		// p为null, 说明一条线走到尽头了
		if !stack.Empty() {
			top := stack.Top().(*Node)
			stack.Pop()
			p = top.Right // 开始打印右节点
		}
	}
}

func LoopInOrder(node *Node) {
	// 中序遍历
	// 1. 往栈加入p,p的左结点赋值给p,不断重复,直到p为空
	// 2.弹出一个元素，访问这个元素， 并将p设置成这个元素的右结点, 重复1
	// 重复1，2直到p和栈都为空
	p := node
	stack := types.NewStack()
	for p != nil || !stack.Empty() {
		for p != nil {
			stack.Push(p)
			p = p.Left
		}
		// p为空
		if !stack.Empty() {
			// 对top来说，左边的元素要么为空， 要么访问过了
			top := stack.Pop().(*Node)
			fmt.Print(top.Name)
			p = top.Right
		}

	}
}

func LoopPostOrder(node *Node) {
	// 1. 给定一个节点p, 加入栈， p设置成p.left, 直到p为空
	// 2. 栈弹出一个元素top
	// 		2.1 如果top->right为空或者top->right等于上一个访问的元素，则访问top, lastNode 设置为 top
	//		2.2 top->right不为空，设置top为p重复 1,2
	stack := types.NewStack()
	p := node
	var lastNode *Node
	for p != nil || !stack.Empty() {
		for p != nil {
			stack.Push(p)
			p = p.Left
		}
		if !stack.Empty() {
			top := stack.Top().(*Node)
			if top.Right == nil || top.Right == lastNode {
				fmt.Print(top.Name)
				lastNode = top
				stack.Pop()
			} else {
				p = top.Right
			}
		}
	}
}

// 根据前序和中序重建二叉树
func RecreateTree(preOrder, inOrder []string) (*Node, error) {
	if preOrder == nil || len(preOrder) != len(inOrder) {
		return nil, errors.New("invalid input")
	}
	return createTree(preOrder, inOrder), nil
}

func createTree(preOrder, inOrder []string) *Node {
	if len(preOrder) == 0 {
		return nil
	}
	rootFlag := preOrder[0]
	root := Node{Name: rootFlag}
	rootPosition := 0
	for i := 0; i < len(inOrder); i++ {
		if inOrder[i] == rootFlag {
			rootPosition = i
		}
	}
	// numbers 表示左子树节点的个数
	numbers := rootPosition
	leftInOrder := inOrder[:rootPosition]
	rightInOrder := inOrder[rootPosition+1:]

	leftPreOrder := preOrder[1 : 1+numbers] //
	rightPreOrder := preOrder[1+numbers:]

	root.Left = createTree(leftPreOrder, leftInOrder)
	root.Right = createTree(rightPreOrder, rightInOrder)

	return &root
}

// 层序遍历
func SequenceOrder(node *Node) {
	qnode := qtypes.NewQueueNode(node)
	currentQueue := qtypes.NewQueue()
	childQueue := qtypes.NewQueue()
	currentQueue.EnQueue(qnode)
	currentQueue.Print()
	for !currentQueue.Empty() || !childQueue.Empty() {
		for !currentQueue.Empty() {
			// 遍历完当前层级Next
			qnode = currentQueue.DeQueue()

			node = qnode.GetData().(*Node)
			fmt.Print(node.Name)
			if node.Left != nil {
				childQueue.EnQueue(qtypes.NewQueueNode(node.Left))
			}
			if node.Right != nil {
				childQueue.EnQueue(qtypes.NewQueueNode(node.Right))
			}
		}
		currentQueue, childQueue = childQueue, currentQueue
	}
}

// 一个带有父节点指针的二叉树，给定一个节点，如何找到中序遍历的下一个节点
/*思路分析
中序遍历的顺序是左->根->右
要找到下一个节点，说明当前节点已经遍历，说明左子树已经被遍历.
1.如果当前节点有右子树，那么下一个输出的节点是右子树最左边的节点。
2.如果当前节点没有有子树，则需要查看他在父节点的位置
	2.1 如果是父节点的左节点, 那下一个输出的就是父节点
	2.2 如果是父节点的右节点，说明父节点已经被输出过了，则要看父节点在父父节点的位置，不断重复
*/

// 前序序列化一个树
func Serializer(head *Node, stream *[]string) {
	var b string
	if head == nil {
		b = "$"
	} else {
		b = head.Name
	}
	*stream = append(*stream, b)
	if head == nil {
		return
	}
	Serializer(head.Left, stream)
	Serializer(head.Right, stream)
}

func GetItem(stream *[]string) func() string {
	i := 0
	var f = func() string {
		res := (*stream)[i]
		i += 1
		return res
	}
	return f
}

func Deserializer(iter func() string, stream *[]string) *Node {
	i := iter()
	if i != "$" {
		node := Node{Name: i}
		left := Deserializer(iter, stream)
		right := Deserializer(iter, stream)
		node.Left = left
		node.Right = right
		return &node
	}
	return nil
}

func RightSpin(node *Node) *Node {
	/*
	     0        0
	    0   ->  0   0
	   0
	*/
	root := node.Left
	node.Left = root.Right
	root.Right = node
	return root
}

func LeftSpin(node *Node) *Node {
	/*
			0            0
		      0     -> 0   0
		        0
	*/
	root := node.Right
	node.Right = root.Left
	root.Left = node
	return root
}
