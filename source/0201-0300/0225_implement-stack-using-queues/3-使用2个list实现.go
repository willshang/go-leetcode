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
/*
入栈过程：
1、q1 为空，放入 q2，否则放入 q1
出栈过程：
1、q1为空：依次取出q2中的元素（除了最后一个），并且放入q1中 取出q2中的最后一个元素，返回结果
否则 依次取出q1中的元素（除了最后一个），并且放入q2中 取出q1中的最后一个元素，返回结果
*/
type MyStack struct {
	l1 *list.List
	l2 *list.List
}

func Constructor() MyStack {
	return MyStack{
		l1: list.New(),
		l2: list.New(),
	}
}

func (m *MyStack) Push(x int) {
	if m.l1.Len() == 0 {
		m.l2.PushBack(x)
	} else {
		m.l1.PushBack(x)
	}
}

func (m *MyStack) Pop() int {
	var top int
	if m.l1.Len() > 0 {
		for m.l1.Len() > 1 {
			m.l2.PushBack(m.l1.Remove(m.l1.Front()))
		}
		top = m.l1.Remove(m.l1.Front()).(int)
	} else {
		for m.l2.Len() > 1 {
			m.l1.PushBack(m.l2.Remove(m.l2.Front()))
		}
		top = m.l2.Remove(m.l2.Front()).(int)
	}
	return top
}

func (m *MyStack) Top() int {
	var top int
	if m.l1.Len() > 0 {
		for m.l1.Len() > 1 {
			m.l2.PushBack(m.l1.Remove(m.l1.Front()))
		}
		top = m.l1.Back().Value.(int)
		m.l2.PushBack(m.l1.Remove(m.l1.Front()))

	} else {
		for m.l2.Len() > 1 {
			m.l1.PushBack(m.l2.Remove(m.l2.Front()))
		}
		top = m.l2.Back().Value.(int)
		m.l1.PushBack(m.l2.Remove(m.l2.Front()))
	}
	return top
}

func (m *MyStack) Empty() bool {
	return m.l1.Len() == 0 && m.l2.Len() == 0
}
