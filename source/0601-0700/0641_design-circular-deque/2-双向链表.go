package main

func main() {

}

type MyCircularDeque struct {
	head   *Node
	tail   *Node
	length int
	cap    int
}

type Node struct {
	value int
	pre   *Node
	next  *Node
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.cap {
		return false
	}
	node := &Node{
		value: value,
	}
	if this.length == 0 {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.pre = node
		this.head = node
	}
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.cap {
		return false
	}
	node := &Node{
		value: value,
	}
	if this.length == 0 {
		this.head = node
		this.tail = node
	} else {
		node.pre = this.tail
		this.tail.next = node
		this.tail = node
	}
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.length == 0 {
		return false
	}
	if this.length == 1 {
		this.head, this.tail = nil, nil
	} else {
		this.head = this.head.next
		this.head.pre = nil
	}
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.length == 0 {
		return false
	}
	if this.length == 1 {
		this.head, this.tail = nil, nil
	} else {
		this.tail = this.tail.pre
		this.tail.next = nil
	}
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}
	return this.head.value
}

func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}
	return this.tail.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.cap
}
