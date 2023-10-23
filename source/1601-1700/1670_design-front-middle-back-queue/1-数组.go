package main

func main() {

}

// leetcode1670_设计前中后队列
type FrontMiddleBackQueue struct {
	arr []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.arr = append([]int{val}, this.arr...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	mid := len(this.arr) / 2
	this.arr = append(this.arr[:mid], append([]int{val}, this.arr[mid:]...)...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.arr = append(this.arr, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	var res int
	if len(this.arr) == 0 {
		return -1
	}
	res = this.arr[0]
	this.arr = this.arr[1:]
	return res
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	var res, mid int
	if len(this.arr) == 0 {
		return -1
	}
	if len(this.arr)%2 == 1 {
		mid = len(this.arr) / 2
	} else {
		mid = len(this.arr)/2 - 1
	}
	res = this.arr[mid]
	this.arr = append(this.arr[:mid], this.arr[mid+1:]...)
	return res
}

func (this *FrontMiddleBackQueue) PopBack() int {
	var res int
	if len(this.arr) == 0 {
		return -1
	}
	res = this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return res
}
