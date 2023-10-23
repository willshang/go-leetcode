package main

import "strings"

func main() {

}

// leetcode677_键值映射
type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{
		m: make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	for key, value := range this.m {
		if strings.HasPrefix(key, prefix) {
			res = res + value
		}
	}
	return res
}
