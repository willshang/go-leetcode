package main

import "sort"

func main() {

}

func maxFrequency(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	res := 1
	total := 0
	i := 0
	for j := 1; j < n; j++ {
		total = total + (nums[j]-nums[j-1])*(j-i) // 累加
		for total > k {
			total = total - (nums[j] - nums[i]) // 不满足，要减去
			i++
		}
		if j-i+1 > res {
			res = j - i + 1
		}
	}
	return res
}
