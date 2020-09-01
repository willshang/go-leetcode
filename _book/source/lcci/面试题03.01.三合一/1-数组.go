package main

func main() {

}

// 程序员面试金典03.01_三合一
type TripleInOne struct {
	arr    []int
	length int
	index  [3]int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		arr:    make([]int, stackSize*3),
		length: stackSize,
		index:  [3]int{0, 0, 0},
	}
}
func (this *TripleInOne) Push(stackNum int, value int) {
	if this.index[stackNum] < this.length {
		this.arr[3*this.index[stackNum]+stackNum] = value
		this.index[stackNum]++
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	res := -1
	if this.index[stackNum] != 0 {
		this.index[stackNum]--
		res = this.arr[3*this.index[stackNum]+stackNum]
	}
	return res
}

func (this *TripleInOne) Peek(stackNum int) int {
	res := -1
	if this.index[stackNum] != 0 {
		res = this.arr[3*(this.index[stackNum]-1)+stackNum]
	}
	return res
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	if this.index[stackNum] == 0 {
		return true
	}
	return false
}
