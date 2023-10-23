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
	nums := make([][]Snap, length)
	for i := 0; i < length; i++ {
		nums[i] = []Snap{{ // 填充0
			id:    0,
			value: 0,
		}}
	}
	return SnapshotArray{
		id:  0,
		arr: nums,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	n := len(this.arr[index])
	if this.arr[index][n-1].id == this.id {
		this.arr[index][n-1].value = val
		return
	}
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
	arr := this.arr[index]
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if snap_id <= arr[mid].id {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left == n || arr[left].id > snap_id {
		return arr[left-1].value
	}
	return arr[left].value
}
