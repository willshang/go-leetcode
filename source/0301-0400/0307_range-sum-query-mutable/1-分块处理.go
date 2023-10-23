package main

import "math"

func main() {

}

type NumArray struct {
	arr    []int // 原数组
	b      []int // 分块和
	length int   // 分块长度
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	per := int(math.Sqrt(float64(n)))
	length := int(math.Ceil(float64(n) / float64(per)))
	b := make([]int, length)
	for i := 0; i < n; i++ {
		b[i/length] = b[i/length] + nums[i]
	}
	return NumArray{
		arr:    nums,
		b:      b,
		length: length,
	}
}

func (this *NumArray) Update(i int, val int) {
	index := i / this.length
	this.b[index] = this.b[index] - this.arr[i] + val
	this.arr[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	res := 0
	a, b := i/this.length, j/this.length
	if a == b {
		for k := i; k <= j; k++ {
			res = res + this.arr[k]
		}
	} else {
		// 分3段
		for k := i; k <= (a+1)*this.length-1; k++ {
			res = res + this.arr[k]
		}
		for k := a + 1; k <= b-1; k++ {
			res = res + this.b[k]
		}
		for k := b * this.length; k <= j; k++ {
			res = res + this.arr[k]
		}
	}
	return res
}
