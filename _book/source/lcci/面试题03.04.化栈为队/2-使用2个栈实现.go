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

// 程序员面试金典03.04_化栈为队
/*
入队: 直接入栈a
出队: 栈b为空，则把栈a中全部数据出栈进入栈b，然后出栈b,不为空直接出栈b
*/
type MyQueue struct {
	a, b *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		a: NewStack(),
		b: NewStack(),
	}
}

func (m *MyQueue) Push(x int) {
	m.a.Push(x)
}

func (m *MyQueue) Pop() int {
	if m.b.Len() == 0 {
		for m.a.Len() > 0 {
			m.b.Push(m.a.Pop())
		}
	}
	return m.b.Pop()
}

func (m *MyQueue) Peek() int {
	res := m.Pop()
	m.b.Push(res)
	return res
}

func (m *MyQueue) Empty() bool {
	return m.a.Len() == 0 && m.b.Len() == 0
}

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{
		nums: []int{},
	}
}

func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
