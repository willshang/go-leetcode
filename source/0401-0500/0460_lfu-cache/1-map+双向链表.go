package main

import "fmt"

func main() {
	cache := Constructor(1)
	cache.Put(2, 1)
	fmt.Println(cache.Get(2)) // 返回 1

	cache.Put(3, 2) // 去除 key 2
	fmt.Println(cache.kv)
	fmt.Println(cache.fk)
	fmt.Println(cache.minFreq)

	fmt.Println(cache.Get(2)) // 返回 1

	fmt.Println(cache.Get(3)) // 返回 1
	//fmt.Println(cache.Get(2)) // 返回 -1 (未找到key 2)
	//fmt.Println(cache.Get(3)) // 返回 3
	//cache.Put(4, 4)           // 去除 key 1
	//fmt.Println(cache.Get(1)) // 返回 -1 (未找到 key 1)
	//fmt.Println(cache.Get(3)) // 返回 3
	//fmt.Println(cache.Get(4)) // 返回 4
}

type Node struct {
	key   int
	value int
	count int
	next  *Node
	prev  *Node
}

type LFUCache struct {
	cap     int
	minFreq int
	kv      map[int]*Node
	fk      map[int]*DoubleList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		kv:      make(map[int]*Node),
		fk:      make(map[int]*DoubleList),
		minFreq: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.kv[key]
	if ok == false {
		return -1
	}
	this.increaseFreq(node)
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	// 已经存在的key，修改并增加次数
	node, ok := this.kv[key]
	if ok {
		node.value = value
		this.increaseFreq(node)
		return
	}
	if this.cap <= len(this.kv) {
		last := this.remove()
		delete(this.kv, last.key)
	}
	temp := &Node{
		key:   key,
		value: value,
		count: 1,
	}
	this.kv[key] = temp
	if this.fk[1] == nil {
		this.fk[1] = NewDoubleList()
	}
	this.fk[1].Push(temp)
	this.minFreq = 1 // 新插入的频率为1
}

func (this *LFUCache) increaseFreq(node *Node) {
	freq := node.count
	this.fk[freq].Remove(node)
	if this.fk[freq+1] == nil {
		this.fk[freq+1] = NewDoubleList()
	}
	node.count++
	this.fk[freq+1].Push(node)
	if this.fk[freq].head.next == this.fk[freq].tail {
		if freq == this.minFreq {
			this.minFreq++
		}
	}
	return
}

func (this *LFUCache) remove() *Node {
	temp := this.fk[this.minFreq]
	last := temp.tail.prev
	temp.Remove(temp.tail.prev) // 删除尾部节点
	return last
}

type DoubleList struct {
	head *Node
	tail *Node
}

func NewDoubleList() *DoubleList {
	node := &DoubleList{
		head: &Node{},
		tail: &Node{},
	}
	node.head.next = node.tail
	node.tail.prev = node.head
	return node
}

// 插入头部
func (this *DoubleList) Push(node *Node) {
	next := this.head.next
	node.next = next
	node.prev = this.head
	next.prev = node
	this.head.next = node
}

// 删除指定节点
func (this *DoubleList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}
