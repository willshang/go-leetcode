package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.arr)
	fmt.Println(obj.SumRange(0, 5))

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	obj = Constructor(nums)
	fmt.Println(obj.arr)
	//[0 1 3 6 10 15 21 28 36 45]
	fmt.Println(obj.SumRange(0, 0))
}

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	size := len(nums)
	arr := make([]int, size+1)
	for i := 1; i <= size; i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	return NumArray{arr: arr}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.arr[j+1] - this.arr[i]
}
