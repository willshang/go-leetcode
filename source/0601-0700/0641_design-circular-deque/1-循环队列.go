package main

func main() {

}

type MyCircularDeque struct {
	arr    []int
	head   int
	tail   int
	length int
}

// leetcode 622.设计循环队列
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		arr:    make([]int, k+1),
		head:   0,
		tail:   0,
		length: k + 1,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = (this.head - 1 + this.length) % this.length // 入队：队头-1
	this.arr[this.head] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr[this.tail] = value
	this.tail = (this.tail + 1 + this.length) % this.length // 入队：队尾+1
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.length // 出队：队头+1
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + this.length) % this.length // 出队：队尾-1
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	index := (this.tail - 1 + this.length) % this.length
	return this.arr[index]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%this.length == this.head
}
