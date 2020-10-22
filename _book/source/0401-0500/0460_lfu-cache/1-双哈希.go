package main

import "container/list"

func main() {

}

type Node struct {
	key   int
	value int
	count int
}

type LFUCache struct {
	cap      int
	dataM    map[int]*list.Element
	count    map[int]*list.List
	minCount int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:      capacity,
		dataM:    make(map[int]*list.Element),
		count:    make(map[int]*list.List),
		minCount: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.dataM[key]
	if ok {
		return this.update(value)
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {

}

func (this *LFUCache) update(value *list.Element) int {
	data := value.Value.(*Node)
	curList, ok := this.count[data.count]
	if !ok {
		return -1
	}
	curList.Remove(ele)
	curList.Remove(value)
	if curList.Len() == 0 {
		if data.count == l.minCount {
			l.minCount++
		}
	}
	data.count++
	newList, ok := l.countMap[data.count]
	if !ok {
		newList = list.New()
		l.countMap[data.count] = newList
	}
	newE := newList.PushBack(data)
	l.findMap[data.key] = newE
	return data.value
}
