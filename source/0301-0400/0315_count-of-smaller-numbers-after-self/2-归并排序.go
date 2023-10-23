package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}

type Node struct {
	Value int
	Index int
}

var res []int

func countSmaller(nums []int) []int {
	n := len(nums)
	res = make([]int, n)
	arr := make([]Node, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, Node{
			Value: nums[i],
			Index: i,
		})
	}
	mergeSort(arr)
	return res
}

func mergeSort(nums []Node) {
	n := len(nums)
	if n <= 1 {
		return
	}
	a := append([]Node{}, nums[:n/2]...)
	b := append([]Node{}, nums[n/2:]...)
	mergeSort(a) // a已经有序
	mergeSort(b) // b已经有序
	i, j := 0, 0
	for i = 0; i < len(a); i++ {
		for j < len(b) && a[i].Value > b[j].Value { // 统计比右边多多少个数
			j++
		}
		res[a[i].Index] = res[a[i].Index] + j // 找到下标，然后加上次数
	}
	i, j = 0, 0
	for k := 0; k < len(nums); k++ { // 2个有序数组合并
		if i < len(a) && (j == len(b) || a[i].Value <= b[j].Value) {
			nums[k] = a[i]
			i++
		} else {
			nums[k] = b[j]
			j++
		}
	}
	return
}
