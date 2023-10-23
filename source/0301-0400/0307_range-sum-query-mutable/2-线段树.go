package main

import "fmt"

func main() {
	node := Constructor([]int{-1})
	fmt.Println(node.SumRange(0, 0))
	node.Update(0, 1)
	fmt.Println(node.SumRange(0, 0))
}

type NumArray struct {
	origin []int // 原数组
	arr    []int // 线段树
	length int   // 长度
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	arr := make([]int, 4*n+1)
	res := NumArray{
		origin: nums,
		arr:    arr,
		length: n,
	}
	for i := 0; i < n; i++ {
		res.UpdateArr(0, 1, res.length, i+1, nums[i]) // 从1开始，添加nums[i]
	}
	return res
}

func (this *NumArray) Update(index int, val int) {
	val, this.origin[index] = val-this.origin[index], val // 需要变更的大小
	this.UpdateArr(0, 1, this.length, index+1, val)       // 从1开始
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Query(0, 1, this.length, left+1, right+1) // 范围+1
}

func (this *NumArray) UpdateArr(id int, left, right, x int, value int) {
	if left > x || right < x {
		return
	}
	this.arr[id] = this.arr[id] + value
	if left == right {
		return
	}
	mid := left + (right-left)/2
	this.UpdateArr(2*id+1, left, mid, x, value)    // 左节点
	this.UpdateArr(2*id+2, mid+1, right, x, value) // 右节点
}

func (this *NumArray) Query(id int, left, right, queryLeft, queryRight int) int {
	if left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		return this.arr[id]
	}
	mid := left + (right-left)/2
	return this.Query(id*2+1, left, mid, queryLeft, queryRight) +
		this.Query(id*2+2, mid+1, right, queryLeft, queryRight)
}
