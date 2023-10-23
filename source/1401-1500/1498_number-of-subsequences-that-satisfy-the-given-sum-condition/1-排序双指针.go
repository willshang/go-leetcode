package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numSubseq([]int{3, 5, 6, 7}, 9))
}

// leetcode1498_满足条件的子序列数目
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	// 计算长度为length满足条件的非空子序列的数目
	// 如1、2、3、4，长度为4，1必选，其他3个数可选可不选，组合数：2^3=8
	m := make(map[int]int)
	m[1] = 1
	for i := 2; i <= len(nums); i++ {
		m[i] = (m[i-1] * 2) % 1000000007
	}
	res := 0
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left]+nums[right] <= target {
			length := right - left + 1
			res = res + m[length]
			left++
		} else {
			right--
		}
	}
	return res % 1000000007
}
