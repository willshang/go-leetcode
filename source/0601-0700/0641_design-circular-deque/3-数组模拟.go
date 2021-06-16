package main

func main() {

}

// leetcode641_设计循环双端队列
type MyCircularDeque struct {
	arr    []int
	length int
	cap    int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		arr:    make([]int, 0),
		length: 0,
		cap:    k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr = append([]int{value}, this.arr...)
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr = append(this.arr, value)
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.arr = this.arr[1:]
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.arr = this.arr[:this.length-1]
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.length-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.cap
}
