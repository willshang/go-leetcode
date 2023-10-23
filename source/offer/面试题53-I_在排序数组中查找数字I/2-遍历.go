package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

// 剑指offer_面试题53-I_在排序数组中查找数字I
func search(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}
