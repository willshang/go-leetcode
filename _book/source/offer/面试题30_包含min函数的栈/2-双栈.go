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

// 剑指offer_面试题30_包含min函数的栈
type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (m *MinStack) Push(x int) {
	if len(m.data) == 0 || x <= m.Min() {
		m.min = append(m.min, x)
	}
	m.data = append(m.data, x)
}

func (m *MinStack) Pop() {
	x := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	if x == m.Min() {
		m.min = m.min[:len(m.min)-1]
	}
}

func (m *MinStack) Top() int {
	if len(m.data) == 0 {
		return 0
	}
	return m.data[len(m.data)-1]
}

func (m *MinStack) Min() int {
	return m.min[len(m.min)-1]
}
