package main

import "fmt"

func main() {
	node := Constructor([]int{1, 3, 5})
	fmt.Println(node.SumRange(0, 2))
}

// leetcode307_区域和检索-数组可修改
type NumArray struct {
	origin []int // 原数组
	c      []int // 树状数组
	length int   // 长度
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	arr := make([]int, n+1)
	res := NumArray{
		origin: nums,
		c:      arr,
		length: n,
	}
	// 单点修改
	for i := 0; i < n; i++ {
		res.UpData(i+1, nums[i]) // 注意下标，默认数组是从1开始
	}
	return res
}

func (this *NumArray) Update(index int, val int) {
	if index < 0 || index > this.length-1 {
		return
	}
	val, this.origin[index] = val-this.origin[index], val // 需要变更的大小
	this.UpData(index+1, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.GetSum(right+1) - this.GetSum(left)
}

func (this *NumArray) LowBit(x int) int {
	return x & (-x)
}

// 单点修改
func (this *NumArray) UpData(i, k int) { // 在i位置加上k
	for i <= this.length {
		this.c[i] = this.c[i] + k
		i = i + this.LowBit(i) // i = i + 2^k
	}
}

// 区间查询
func (this *NumArray) GetSum(i int) int {
	res := 0
	for i > 0 {
		res = res + this.c[i]
		i = i - this.LowBit(i)
	}
	return res
}
