package main

func main() {
	
}
type MyCircularDeque struct {
	f, r *node
	len, cap int
}

type node struct {
	value int
	pre, next *node
}



/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap:k,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.len == this.cap{
		return false
	}
	n := &node{
		value:value,
	}

	if this.len == 0{
		this.f = n
		this.r = n
	}else {
		n.next = this.f
		this.f.pre = n
		this.f = n
	}
	this.len++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.len == this.cap{
		return false
	}
	n := &node{
		value:value,
	}

	if this.len == 0{
		this.f = n
		this.r = n
	}else {
		n.pre = this.r
		this.r.next = n
		this.r = n
	}
	this.len++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.len == 0{
		return false
	}
	if this.len == 1{
		this.f, this.r = nil,nil
	}else {
		this.f = this.f.next
		this.f.pre = nil
	}
	this.len--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.len == 0{
		return false
	}
	if this.len == 1{
		this.f, this.r = nil, nil
	}else {
		this.r = this.r.pre
		this.r.next = nil
	}
	this.len--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.len == 0{
		return -1
	}
	return this.f.value
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.len == 0{
		return -1
	}
	return this.r.value
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.cap
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */