package main

import "fmt"

func main() {
	obj := Constructor()
	x := 5
	obj.Push(x)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}

// leetcode232_用栈实现队列
type MyQueue struct {
	a []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (m *MyQueue) Push(x int) {
	m.a = append(m.a, x)
}

func (m *MyQueue) Pop() int {
	if len(m.a) == 0 {
		return 0
	}
	first := m.a[0]
	m.a = m.a[1:]
	return first
}

func (m *MyQueue) Peek() int {
	if len(m.a) == 0 {
		return 0
	}
	return m.a[0]
}

func (m *MyQueue) Empty() bool {
	if len(m.a) == 0 {
		return true
	}
	return false
}
