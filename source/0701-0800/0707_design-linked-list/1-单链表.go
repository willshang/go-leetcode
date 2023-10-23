package main

func main() {

}

type MyLinkedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
}

func Constructor() MyLinkedList {
	tail := &Node{}
	head := &Node{next: tail}
	return MyLinkedList{
		head: head,
		tail: tail,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.size <= index {
		return -1
	}
	i := 0
	cur := this.head.next
	for i < index {
		i++
		cur = cur.next
	}
	return cur.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{value: val}
	node.next = this.head.next
	this.head.next = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.tail.value = val
	node := &Node{}
	this.tail.next = node
	this.tail = node
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index < 0 || this.size < index:
		return
	case index == 0:
		this.AddAtHead(val)
		return
	case index == this.size:
		this.AddAtTail(val)
		return
	}
	i := -1
	cur := this.head
	for i+1 < index {
		i++
		cur = cur.next
	}
	node := &Node{value: val}
	node.next = cur.next
	cur.next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.size <= index {
		return
	}
	i := -1
	cur := this.head
	for i+1 < index {
		i++
		cur = cur.next
	}
	cur.next = cur.next.next
	this.size--
}
