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
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			count := 1
			for j := i + 1; j < n; j++ {
				if nums[j] > nums[i]+count {
					break
				}
				if nums[j] >= 0 && nums[j] == nums[i]+count {
					nums[j] = -1
					count++
					if count == k {
						break
					}
				}
			}
			if count != k {
				return false
			}
			nums[i] = -1
		}
	}
	return true
}
