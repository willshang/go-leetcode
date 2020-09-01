package main

func main() {

}

// 程序员面试金典03.03_堆盘子
type StackOfPlates struct {
	cap   int
	stack [][]int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		cap:   cap,
		stack: make([][]int, 0),
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	if len(this.stack) == 0 {
		newStack := make([]int, 0)
		newStack = append(newStack, val)
		this.stack = append(this.stack, newStack)
		return
	}
	last := this.stack[len(this.stack)-1]
	if len(last) == this.cap {
		newStack := make([]int, 0)
		newStack = append(newStack, val)
		this.stack = append(this.stack, newStack)
		return
	}
	last = append(last, val)
	this.stack[len(this.stack)-1] = last
}

func (this *StackOfPlates) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	last := this.stack[len(this.stack)-1]
	res := last[len(last)-1]
	last = last[:len(last)-1]
	this.stack[len(this.stack)-1] = last
	if len(last) == 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
	return res
}

func (this *StackOfPlates) PopAt(index int) int {
	if index >= len(this.stack) {
		return -1
	}
	arr := this.stack[index]
	res := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	this.stack[index] = arr
	if len(arr) == 0 {
		this.stack = append(this.stack[:index], this.stack[index+1:]...)
	}
	return res
}
