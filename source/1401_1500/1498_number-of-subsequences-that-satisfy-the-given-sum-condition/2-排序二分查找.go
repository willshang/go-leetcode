package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numSubseq([]int{3, 5, 6, 7}, 9))
}

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
	for i := 0; i < len(nums); i++ {
		left, right := i, len(nums)
		for left+1 < right {
			mid := left + (right-left)/2
			if nums[mid]+nums[i] <= target {
				left = mid
			} else {
				right = mid
			}
		}
		if nums[left]+nums[i] <= target {
			length := left - i + 1
			res = res + m[length]
		}
	}
	return res % 1000000007
}