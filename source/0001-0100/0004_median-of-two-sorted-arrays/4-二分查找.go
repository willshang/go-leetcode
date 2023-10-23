package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}
	left, right := 0, n
	a, b := 0, 0
	for left <= right {
		// 左半部分最大的值小于等于右半部分最小的值: max(A[i-1],B[j-1])) <= min(A[i],B[j]))
		i := left + (right-left)/2 // i,j分别对num1,num2的划分
		j := (n+m+1)/2 - i         // i+j == (n+m+1)/2
		// 偶数求a=>max(A[i-1],B[j-1])) b=>min(A[i],B[j]))
		// 奇数求a=>max(A[i-1],B[j-1]))
		if j != 0 && i != n && nums1[i] < nums2[j-1] {
			left = i + 1
		} else if i != 0 && j != m && nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			if i == 0 {
				a = nums2[j-1]
			} else if j == 0 {
				a = nums1[i-1]
			} else {
				a = max(nums1[i-1], nums2[j-1])
			}
			if (n+m)%2 == 1 {
				return float64(a)
			}
			if i == n {
				b = nums2[j]
			} else if j == m {
				b = nums1[i]
			} else {
				b = min(nums1[i], nums2[j])
			}
			return float64(a+b) / 2.0
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
