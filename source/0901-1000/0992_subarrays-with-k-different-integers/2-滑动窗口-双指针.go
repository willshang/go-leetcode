package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
}

func subarraysWithKDistinct(nums []int, k int) int {
	res := 0
	n := len(nums)
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	left1, left2 := 0, 0
	count1, count2 := 0, 0
	for i := 0; i < n; i++ {
		cur := nums[i]
		if m1[cur] == 0 {
			count1++
		}
		if m2[cur] == 0 {
			count2++
		}
		m1[cur]++
		m2[cur]++
		for count1 > k {
			m1[nums[left1]]--
			if m1[nums[left1]] == 0 {
				count1--
			}
			left1++
		}
		for count2 > k-1 {
			m2[nums[left2]]--
			if m2[nums[left2]] == 0 {
				count2--
			}
			left2++
		}
		res = res + left2 - left1
	}
	return res
}
