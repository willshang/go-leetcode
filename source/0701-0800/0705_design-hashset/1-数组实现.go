package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Add(6)
	obj.Add(7)
	obj.Remove(6)
	fmt.Println(obj.table)
}

// leetcode705_设计哈希集合
type MyHashSet struct {
	table []bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		table: make([]bool, 1000001),
	}
}

func (m *MyHashSet) Add(key int) {
	m.table[key] = true
}

func (m *MyHashSet) Remove(key int) {
	m.table[key] = false
}

func (m *MyHashSet) Contains(key int) bool {
	return m.table[key]
}
