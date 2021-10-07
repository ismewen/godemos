package types

import (
	"fmt"
	"sync"
)

type StackItem interface{}

type StackInterface interface {
	Pop() StackItem
	Push(item StackItem)
	Len() int
	Empty() bool
	Top() StackItem
}

type Stack struct {
	items []StackItem
	sync.RWMutex
}

func (stack *Stack) Len() int {
	return len(stack.items)
}

func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Pop() StackItem {
	stack.Lock()
	defer stack.Unlock()

	pop := stack.items[stack.Len()-1]
	stack.items = stack.items[:stack.Len()-1]

	return pop
}

func (stack *Stack) Push(item StackItem) {
	stack.Lock()
	defer stack.Unlock()

	stack.items = append(stack.items, item)
}

func (stack *Stack) Top() StackItem {
	stack.RLock()
	defer stack.RUnlock()
	if !stack.Empty() {
		return stack.items[stack.Len()-1]
	}
	return nil
}

func (stack *Stack) Print() {
	for _, v := range stack.items {
		fmt.Println(v, " ")
	}
}

func NewStack() *Stack {
	stack := Stack{
		items: []StackItem{},
	}
	return &stack
}

type MStack struct {
	DataStack *Stack
	MinStack *Stack
}

func (ms *MStack) Push(){

}

func (ms *MStack) Pop(){

}


