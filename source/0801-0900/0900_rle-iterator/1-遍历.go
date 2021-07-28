package main

func main() {

}

// leetcode900_RLE迭代器
type RLEIterator struct {
	arr   []int
	index int // 第几个元素
	count int // 元素的第几次
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{
		arr:   encoding,
		index: 0,
		count: 0,
	}
}

func (this *RLEIterator) Next(n int) int {
	for this.index < len(this.arr) {
		if n+this.count > this.arr[this.index] {
			n = n - (this.arr[this.index] - this.count)
			this.count = 0
			this.index = this.index + 2
		} else {
			this.count = this.count + n
			return this.arr[this.index+1]
		}
	}
	return -1
}
