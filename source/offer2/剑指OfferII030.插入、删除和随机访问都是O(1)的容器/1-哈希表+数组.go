package main

import "math/rand"

func main() {

}

// 剑指OfferII030.插入、删除和随机访问都是O(1)的容器
type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:   make(map[int]int),
		arr: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.m[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	index := this.m[val]
	this.arr[index], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[index]
	this.m[this.arr[index]] = index
	this.arr = this.arr[:len(this.arr)-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.arr) == 0 {
		return -1
	}
	index := rand.Intn(len(this.arr))
	return this.arr[index]
}
