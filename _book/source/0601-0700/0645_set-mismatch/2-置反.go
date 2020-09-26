package main

import "fmt"

func main() {
	nums := []int{2, 1, 5, 2, 4}
	fmt.Println(findErrorNums(nums))
}

// leetcode645_错误的集合
func findErrorNums(nums []int) []int {
	repeatNum := 0
	for i := 0; i < len(nums); i++ {
		n := abs(nums[i])
		if nums[n-1] < 0 {
			repeatNum = n
		} else {
			nums[n-1] = -nums[n-1]
		}
	}
	misNum := 0
	for i, v := range nums {
		if v > 0 {
			misNum = i + 1
			break
		}
	}
	return []int{repeatNum, misNum}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
