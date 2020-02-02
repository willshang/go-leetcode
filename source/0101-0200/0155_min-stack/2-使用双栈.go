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

type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.data) == 0 || x <= this.GetMin() {
		this.min = append(this.min, x)
	}
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	x := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if x == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
