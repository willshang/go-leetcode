package main

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Pop()

}

// 程序员面试金典03.05_栈排序
type SortedStack struct {
	stack []int
	temp  []int
}

func Constructor() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int) {
	for len(this.stack) > 0 && val >= this.stack[len(this.stack)-1] {
		this.temp = append(this.temp, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, val)
	for len(this.temp) > 0 {
		this.stack = append(this.stack, this.temp[len(this.temp)-1])
		this.temp = this.temp[:len(this.temp)-1]
	}
}

func (this *SortedStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *SortedStack) Peek() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.stack) == 0
}
