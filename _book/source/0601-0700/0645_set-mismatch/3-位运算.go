package main

import "fmt"

func main() {
	//nums := []int{3, 1, 7, 6, 4,2,7}
	//nums := []int{1,2,2,4}
	nums := []int{1, 1}
	fmt.Println(findErrorNums(nums))
}

func findErrorNums(nums []int) []int {
	res := 0
	// 异或得到repeatedNum^misNum
	for i := 0; i < len(nums); i++ {
		res = res ^ (i + 1) ^ (nums[i])
	}
	// 找到第一位不是0的
	h := 1
	for res&h == 0 {
		h = h << 1
	}
	a := 0
	b := 0
	for i := range nums {
		if h&nums[i] == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
		if h&(i+1) == 0 {
			a ^= i + 1
		} else {
			b ^= i + 1
		}
	}
	for i := range nums {
		if nums[i] == b {
			return []int{b, a}
		}
	}
	return []int{a, b}
}
