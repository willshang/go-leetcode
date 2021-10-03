package main

import "sort"

func main() {

}

// 剑指OfferII061.和最小的k个数对
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	arr := make([][]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			arr = append(arr, []int{nums1[i], nums2[j]})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0]+arr[i][1] < arr[j][0]+arr[j][1]
	})
	if len(arr) < k {
		return arr
	}
	return arr[:k]
}
