package main

import "container/list"

func main() {

}

type MaxQueue struct {
	data *list.List
	max  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		data: list.New(),
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
	this.data.PushBack(value)
	for len(this.max) > 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	res := -1
	if this.data.Len() > 0 {
		res = this.data.Front().Value.(int)
		this.data.Remove(this.data.Front())
		if res == this.max[0] {
			this.max = this.max[1:]
		}
	}
	return res
}
