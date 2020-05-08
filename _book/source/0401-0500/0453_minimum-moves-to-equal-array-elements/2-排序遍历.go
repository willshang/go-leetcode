package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 10}
	fmt.Println(minMoves(arr))
}

// n-1个元素+1等价于1个元素-1，将问题转化为：对所有元素到最小元素的差进行求和

// leetcode453_最小移动次数使数组元素相等
func minMoves(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		sum = sum + nums[i] - nums[0]
	}
	return sum
}
