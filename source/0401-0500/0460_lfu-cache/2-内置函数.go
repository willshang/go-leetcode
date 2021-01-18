package main

import "container/list"

// leetcode460_LFU缓存
type Node struct {
	key   int
	value int
	count int
}

type LFUCache struct {
	cap     int
	minFreq int
	kv      map[int]*list.Element
	fk      map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		minFreq: 0,
		kv:      make(map[int]*list.Element),
		fk:      make(map[int]*list.List),
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.kv[key]
	if ok == false {
		return -1
	}
	return this.increaseFreq(node)
}

func (this *LFUCache) Put(key int, value int) {
	data, ok := this.kv[key]
	if ok {
		node := data.Value.(*Node)
		node.value = value
		this.increaseFreq(data)
		return
	}
	if this.cap == len(this.kv) {
		cur, ok := this.fk[this.minFreq]
		if ok == false {
			return
		}
		deleteKey := cur.Front()
		cur.Remove(deleteKey)
		delete(this.kv, deleteKey.Value.(*Node).key)
	}
	temp := &Node{
		key:   key,
		value: value,
		count: 1,
	}
	if _, ok := this.fk[1]; ok == false {
		this.fk[1] = list.New()
	}
	res := this.fk[1].PushBack(temp)
	this.kv[key] = res
	this.minFreq = 1
}

func (this *LFUCache) increaseFreq(data *list.Element) int {
	node := data.Value.(*Node)
	cur, ok := this.fk[node.count]
	if ok == false {
		return -1
	}
	cur.Remove(data)
	if cur.Len() == 0 && this.minFreq == node.count {
		this.minFreq++
	}
	node.count++
	if this.fk[node.count] == nil {
		this.fk[node.count] = list.New()
	}
	res := this.fk[node.count].PushBack(node)
	this.kv[node.key] = res
	return node.value
}
