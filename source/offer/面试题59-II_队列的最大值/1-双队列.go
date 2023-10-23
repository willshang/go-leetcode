package main

func main() {

}

// 剑指offer_面试题59-II_队列的最大值
type MaxQueue struct {
	data []int
	max  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		data: make([]int, 0),
		max:  make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	for len(this.max) > 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	res := -1
	if len(this.data) > 0 {
		res = this.data[0]
		this.data = this.data[1:]
		if res == this.max[0] {
			this.max = this.max[1:]
		}
	}
	return res
}
