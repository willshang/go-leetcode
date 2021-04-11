package main

import "sort"

func main() {

}

// leetcode1818_绝对差值和
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	arr := make([]int, len(nums1))
	sum := 0
	for i := 0; i < len(nums1); i++ {
		arr[i] = nums1[i]
		sum = (sum + abs(nums1[i]-nums2[i])) % 1000000007
	}
	sort.Ints(arr)
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if nums1[i] == nums2[i] {
			continue
		}
		b := nums2[i]
		target := search(arr, b)
		maxValue = max(maxValue, abs(nums1[i]-b)-abs(target-b))
	}
	return (sum - maxValue + 1000000007) % 1000000007
}

func search(arr []int, target int) int {
	res := 0
	if arr[0] > target {
		return arr[0]
	}
	if arr[len(arr)-1] < target {
		return arr[len(arr)-1]
	}
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if target < arr[mid] {
			right = mid - 1
			if abs(res-target) > abs(target-arr[mid]) {
				res = arr[mid]
			}
		} else if target == arr[mid] {
			return target
		} else if target > arr[mid] {
			left = mid + 1
			if abs(res-target) > abs(target-arr[mid]) {
				res = arr[mid]
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
