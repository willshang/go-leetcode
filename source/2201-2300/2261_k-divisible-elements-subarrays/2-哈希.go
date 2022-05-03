package main

import "fmt"

func main() {
	fmt.Println(countDistinct([]int{2, 3, 3, 2, 2}, 2, 2))
}

// leetcode2261_含最多K个可整除元素的子数组
func countDistinct(nums []int, k int, p int) int {
	m := make(map[[200]int]bool)
	n := len(nums)
	for i := 0; i < n; i++ { // 左边界
		count := 0
		temp := [200]int{}
		for j := i; j < n; j++ { // 右边界
			if nums[j]%p == 0 {
				count++
			}
			temp[j-i] = nums[j]
			if count <= k {
				m[temp] = true
			}
		}
	}
	return len(m)
}
