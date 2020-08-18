package main

import "fmt"

func main() {
	node := Constructor(3)
	node.Put(2, 1)
	fmt.Println(node)
}

// 程序员面试金典16.25_LRU缓存
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.prev = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.remove(node)
		this.putHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value
		this.remove(node)
		this.putHead(node)
		return
	}
	if this.cap <= len(this.m) {
		// 删除尾部
		deleteKey := this.tail.prev.key
		this.remove(this.tail.prev)
		delete(this.m, deleteKey)
	}
	// 插入到头部
	newNode := &Node{key: key, value: value}
	this.putHead(newNode)
	this.m[key] = newNode
}

// 删除尾部节点
func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 插入头部
func (this *LRUCache) putHead(node *Node) {
	next := this.header.next
	this.header.next = node
	node.next = next
	next.prev = node
	node.prev = this.header
}
