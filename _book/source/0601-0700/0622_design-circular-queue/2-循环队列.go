package main

func main() {

}

type MyCircularQueue struct {
	queue []int
	k     int
	front int // 队首
	rear  int // 队尾
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k+1),
		k:     k + 1,
		front: 0,
		rear:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	// 队尾入队
	this.queue[this.rear] = value
	this.rear++
	if this.rear == this.k {
		this.rear = 0
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	// 队尾出队
	this.front++
	if this.front == this.k {
		this.front = 0
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	prev := this.rear - 1
	if prev < 0 {
		prev = this.k - 1
	}
	return this.queue[prev]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularQueue) IsFull() bool {
	next := this.rear + 1
	if next == this.k {
		next = 0
	}
	return next == this.front
}
