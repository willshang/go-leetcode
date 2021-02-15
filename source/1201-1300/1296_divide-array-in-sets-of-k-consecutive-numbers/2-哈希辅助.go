package main

import "sort"

func main() {

}

func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}
	if k == 1 {
		return true
	}
	arr := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 0 {
			arr = append(arr, nums[i])
		}
		m[nums[i]]++
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if m[arr[i]] > 0 {
			for j := 1; j < k; j++ {
				value := arr[i] + j
				m[value] = m[value] - m[arr[i]]
				if m[value] < 0 {
					return false
				}
			}
		}
	}
	return true
}
