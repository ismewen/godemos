package types

import (
	"errors"
	"fmt"
	"sync"
)

const LINKLISTMAXSIZE = 255

// 单链表接口定义

type SingleLinkList interface {
	Insert(i int, node *Node) error
	Delete(i int) (*Node, error)
}

// 链表存储实现
type Node struct {
	Data interface{}
	Next *Node
}

func (node *Node) GetData() interface{} {
	return node.Data
}

type Head struct {
	Name   string
	Next   *Node
	length int
	sync.RWMutex
}

func NewHead(name string ) *Head {
	h := Head{
		Name: name,
	}
	return &h
}

func (head *Head) Len() int {
	return head.length
}

func (head *Head) Insert(i int, node *Node) error {
	head.Lock()
	defer head.Unlock()

	// head 和 node 数据结构不一样，需要单独处理
	if i == 0 {
		head.Next, node.Next = node, head.Next
	}

	if i > head.length {
		return errors.New("")
	}
	preNode := head.Next

	for j := 0; j < i-1; j++ {
		preNode = preNode.Next
	}
	head.length += 1
	preNode.Next, node.Next = node, preNode.Next
	return nil
}

func (head *Head) Delete(i int) (*Node, error) {
	if i >= head.length {
		return nil, errors.New("IndexError")
	}
	if i == 0 {
		deleteNode := head.Next
		head.Next = deleteNode.Next
		head.length -= 1
		return deleteNode, nil
	}
	preNode := head.Next
	for j := 0; j < i-1; j++ {
		preNode = preNode.Next
	}
	deleteNode := preNode.Next
	preNode.Next = preNode.Next.Next
	head.length -= 1
	return deleteNode, nil
}

func (head *Head) Print() {
	cur := head.Next
	for cur != nil {
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}
func (head *Head) Clear() {
	// 清空
	head.Lock()
	defer head.Unlock()

	head.length = 0
	head.Next = nil
}

type LinkItem interface{}

// 数组存储实现
type LinkList struct {
	Name   string
	items  [LINKLISTMAXSIZE]LinkItem
	length int
	sync.RWMutex
}

func (ll *LinkList) Insert(i int, item LinkItem) error {
	ll.Lock()
	defer ll.Unlock()

	if i > ll.length {
		return errors.New("index error")
	}

	if ll.length == LINKLISTMAXSIZE {
		// 队列已满
		return errors.New("full Error")
	}

	for j := ll.length; j > i; j-- {
		ll.items[j] = ll.items[j-1]
	}

	ll.items[i] = item

	ll.length += 1
	return nil
}

func (ll *LinkList) Delete(i int) (LinkItem, error) {
	ll.Lock()
	defer ll.Unlock()

	if i >= ll.length {
		return nil, errors.New("index error")
	}
	deleteItem := ll.items[i]
	for j := i; j < ll.length-1; j++ {
		ll.items[j] = ll.items[j+1]
	}
	ll.length -= 1
	return deleteItem, nil
}

func (ll *LinkList) NewIterator() Iterator {
	return Iterator{cur: 0, LinkList: ll}
}

type Iterator struct {
	cur int
	*LinkList
}

func (it *Iterator) HasNext() bool {
	return it.cur >= it.length
}

func (it *Iterator) Next() LinkItem {
	it.cur += 1
	return it.items[it.cur]
}
