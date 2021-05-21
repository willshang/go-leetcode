package main

import "fmt"

func main() {
	a := Constructor(1)
	a.Set(0, 15)
	fmt.Println(a.Snap())
	fmt.Println(a.Snap())
	fmt.Println(a.Snap())

	fmt.Println("get ", a.Get(0, 2))
	fmt.Println(a.Snap())
	fmt.Println(a.Snap())
	fmt.Println(a)
	fmt.Println("get ", a.Get(0, 0))
}

// leetcode1146_快照数组
type SnapshotArray struct {
	id  int
	arr []map[int]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		id:  0,
		arr: make([]map[int]int, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	if this.arr[index] == nil {
		this.arr[index] = make(map[int]int)
	}
	this.arr[index][this.id] = val
}

func (this *SnapshotArray) Snap() int {
	id := this.id
	this.id++
	return id
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	if this.arr[index] == nil {
		return 0
	}
	for ; snap_id >= 0; snap_id-- {
		if v, ok := this.arr[index][snap_id]; ok {
			return v
		}
	}
	return 0
}
