package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Min()

	fmt.Println(param_3, param_4)
}

type item struct {
	min, x int
}

type MinStack struct {
	stack []item
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	min := x
	if len(m.stack) > 0 && m.Min() < x {
		min = m.Min()
	}
	m.stack = append(m.stack, item{
		min: min,
		x:   x,
	})
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) Top() int {
	if len(m.stack) == 0 {
		return 0
	}
	return m.stack[len(m.stack)-1].x
}

func (m *MinStack) Min() int {
	if len(m.stack) == 0 {
		return 0
	}
	return m.stack[len(m.stack)-1].min
}
