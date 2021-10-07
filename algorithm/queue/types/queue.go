package types

import (
	"algorithm/linklist/types"
)

type QueueItem interface{}

type Queue struct {
	*types.Head // 通过单链表实现
}

type QueueInterface interface {
	EnQueue()
	DeQueue()
	Top()
	IsEmpty() bool
}

func (queue *Queue) EnQueue(node *types.Node) {
	_ = queue.Insert(queue.Len(), node)
}

func (queue *Queue) DeQueue() *types.Node {
	node, _ := queue.Delete(0)
	return node
}

func (queue *Queue) Top() *types.Node {
	return queue.Next
}

func (queue *Queue) Empty() bool {
	return queue.Len() == 0
}

func NewQueue() *Queue {
	return &Queue{Head: &types.Head{}}
}

func NewQueueNode(data interface{}) *types.Node {
	return &types.Node{Data: data}
}
