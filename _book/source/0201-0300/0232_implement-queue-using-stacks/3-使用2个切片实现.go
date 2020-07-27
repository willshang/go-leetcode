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

type MyQueue struct {
	a []int
	b []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (m *MyQueue) Push(x int) {
	m.a = append(m.a, x)
}

func (m *MyQueue) Pop() int {
	m.Peek()
	temp := m.b[len(m.b)-1]
	m.b = m.b[:len(m.b)-1]
	return temp
}

func (m *MyQueue) Peek() int {
	if len(m.b) == 0 {
		for len(m.a) > 0 {
			m.b = append(m.b, m.a[len(m.a)-1])
			m.a = m.a[:len(m.a)-1]
		}
	}
	if len(m.b) == 0 {
		return -1
	}
	return m.b[len(m.b)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.a) == 0 && len(m.b) == 0
}
