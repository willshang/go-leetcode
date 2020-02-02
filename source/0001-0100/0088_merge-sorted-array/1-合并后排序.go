package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{}, 0)
	fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2[:n]...)
	sort.Ints(nums1)
}
