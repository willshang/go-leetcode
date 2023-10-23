package main

import "sort"

func main() {

}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	arr := make([]int, nums[n-1]-nums[0]+1)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			arr[nums[j]-nums[i]]++
		}
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		count = count + arr[i]
		if count >= k {
			return i
		}
	}
	return 0
}
