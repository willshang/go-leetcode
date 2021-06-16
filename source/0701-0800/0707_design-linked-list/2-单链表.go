package main

func main() {

}

// leetcode707_设计链表
type MyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &Node{},
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	prev := this.head
	for i := 1; i <= index; i++ {
		prev = prev.next
	}
	return prev.next.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	prev := this.head
	for i := 1; i <= index; i++ {
		prev = prev.next
	}
	node := &Node{
		value: val,
		next:  nil,
	}
	node.next = prev.next
	prev.next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.head
	for i := 1; i <= index; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	this.size--
}
