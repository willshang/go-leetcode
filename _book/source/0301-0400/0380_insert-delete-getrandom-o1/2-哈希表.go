package main

import (
	"math/rand"
)

func main() {

}

type RandomizedSet struct {
	m map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.m) == 0 {
		return -1
	}
	index := rand.Intn(len(this.m))
	res := -1
	for res = range this.m {
		index--
		if index == -1 {
			break
		}
	}
	return res
}
