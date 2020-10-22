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

type MyStack struct {
	l1 *Queue
	l2 *Queue
}

func Constructor() MyStack {
	return MyStack{
		l1: NewQueue(),
		l2: NewQueue(),
	}
}

func (m *MyStack) Push(x int) {
	m.l1.Push(x)
}

func (m *MyStack) Pop() int {
	if m.l2.Len() == 0 {
		m.l1, m.l2 = m.l2, m.l1
	}

	for m.l2.Len() > 1 {
		m.l1.Push(m.l2.Pop())
	}
	return m.l2.Pop()
}

func (m *MyStack) Top() int {
	res := m.Pop()
	m.l1.Push(res)
	return res
}

func (m *MyStack) Empty() bool {
	return (m.l1.Len() + m.l2.Len()) == 0
}

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{
		nums: []int{},
	}
}

func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) Pop() int {
	if len(q.nums) == 0 {
		return 0
	}
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
