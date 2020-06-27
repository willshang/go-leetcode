package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Add(6)
	obj.Add(7)
	obj.Remove(6)
	fmt.Println(obj.table)
}

type MyHashSet struct {
	table [10000][]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		table: [10000][]int{},
	}
}

func (m *MyHashSet) Add(key int) {
	for _, v := range m.table[key%10000] {
		if v == key {
			return
		}
	}
	m.table[key%10000] = append(m.table[key%10000], key)
}

func (m *MyHashSet) Remove(key int) {
	for k, v := range m.table[key%10000] {
		if v == key {
			m.table[key%10000] = append(m.table[key%10000][:k], m.table[key%10000][k+1:]...)
		}
	}
}

func (m *MyHashSet) Contains(key int) bool {
	for _, v := range m.table[key%10000] {
		if v == key {
			return true
		}
	}
	return false
}
