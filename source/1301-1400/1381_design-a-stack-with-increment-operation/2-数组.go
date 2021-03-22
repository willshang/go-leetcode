package main

import "math"

func main() {

}

type CustomStack struct {
	stack []int
	add   []int
	top   int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, maxSize),
		add:   make([]int, maxSize),
		top:   -1,
	}
}

func (this *CustomStack) Push(x int) {
	if this.top != len(this.stack)-1 {
		this.top++
		this.stack[this.top] = x
	}
}

func (this *CustomStack) Pop() int {
	if this.top == -1 {
		return -1
	}
	res := this.stack[this.top] + this.add[this.top]
	if this.top != 0 {
		this.add[this.top-1] = this.add[this.top-1] + this.add[this.top]
	}
	this.add[this.top] = 0
	this.top--
	return res
}

func (this *CustomStack) Increment(k int, val int) {
	index := int(math.Min(float64(k-1), float64(this.top)))
	if index >= 0 {
		this.add[index] = this.add[index] + val
	}
}
