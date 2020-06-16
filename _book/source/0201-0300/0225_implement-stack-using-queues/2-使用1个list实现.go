package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	x := 5
	obj.Push(x)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()

	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}

// leetcode225_用队列实现栈
type MyStack struct {
	*list.List
}

func Constructor() MyStack {
	return MyStack{
		list.New(),
	}
}

func (m *MyStack) Push(x int) {
	m.PushBack(x)
}

func (m *MyStack) Pop() int {
	if m.Len() == 0 {
		return -1
	}
	return m.Remove(m.Back()).(int)
}

func (m *MyStack) Top() int {
	if m.Len() == 0 {
		return -1
	}
	return m.Back().Value.(int)
}

func (m *MyStack) Empty() bool {
	return m.Len() == 0
}
