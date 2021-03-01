package main

func main() {

}

// leetcode1381_设计一个支持增量操作的栈
type CustomStack struct {
	stack []int
	size  int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, 0),
		size:  maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) < this.size {
		this.stack = append(this.stack, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.stack) > 0 {
		res := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return res
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	if k > len(this.stack) {
		k = len(this.stack)
	}
	for i := 0; i < k; i++ {
		this.stack[i] = this.stack[i] + val
	}
}
