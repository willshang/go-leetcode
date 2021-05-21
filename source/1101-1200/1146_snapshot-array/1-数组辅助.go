package main

import "fmt"

func main() {
	a := Constructor(1)
	a.Set(0, 15)
	fmt.Println(a.Snap())
	fmt.Println(a.Snap())
	fmt.Println(a.Snap())

	fmt.Println(a.Get(0, 2))
	fmt.Println(a.Snap())
	fmt.Println(a.Snap())
	fmt.Println(a)
	a.Get(0, 0)
}

type SnapshotArray struct {
	id  int
	arr [][]Snap
}

type Snap struct {
	id    int
	value int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		id:  0,
		arr: make([][]Snap, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	// 保存的是index的操作记录
	this.arr[index] = append(this.arr[index], Snap{
		id:    this.id,
		value: val,
	})
}

func (this *SnapshotArray) Snap() int {
	id := this.id
	this.id++
	return id
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	n := len(this.arr[index])
	i := 0
	// 根据index的历史记录，找到最后1个id为snap_id的记录
	// 其中id为snap_id的记录可能有多个
	for i < n && this.arr[index][i].id <= snap_id {
		i++
	}
	if i == 0 {
		return 0
	}
	i--
	return this.arr[index][i].value
}
