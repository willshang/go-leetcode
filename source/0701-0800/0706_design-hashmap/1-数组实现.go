package main

func main() {

}

// leetcode706_设计哈希映射
type MyHashMap struct {
	table []int
}

func Constructor() MyHashMap {
	return MyHashMap{
		table: make([]int, 1000001),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.table[key] = value + 1
}

func (this *MyHashMap) Get(key int) int {
	return this.table[key] - 1
}

func (this *MyHashMap) Remove(key int) {
	this.table[key] = 0
}
