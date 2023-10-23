package main

import "fmt"

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
	arr []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(x int) {
	m.arr = append(m.arr, x)
}

func (m *MyStack) Pop() int {
	if len(m.arr) == 0 {
		return 0
	}
	last := m.arr[len(m.arr)-1]
	m.arr = m.arr[0 : len(m.arr)-1]
	return last
}

func (m *MyStack) Top() int {
	if len(m.arr) == 0 {
		return 0
	}
	return m.arr[len(m.arr)-1]
}

func (m *MyStack) Empty() bool {
	if len(m.arr) == 0 {
		return true
	}
	return false
}
