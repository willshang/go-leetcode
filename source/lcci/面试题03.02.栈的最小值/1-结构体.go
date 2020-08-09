package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Println(param_3, param_4)
}

// 程序员面试金典03.02_栈的最小值
type item struct {
	min, x int
}

type MinStack struct {
	stack []item
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{
		min: min,
		x:   x,
	})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].x
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].min
}
