package main

func main() {

}

// leetcode622_设计循环队列
type MyCircularQueue struct {
	queue []int
	k     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, 0),
		k:     k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.queue) == this.k {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.queue) == 0 {
		return false
	}
	this.queue = this.queue[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

func (this *MyCircularQueue) Rear() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.queue) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.queue) == this.k
}
