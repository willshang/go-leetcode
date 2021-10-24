package main

import "math/bits"

func main() {

}

// leetcode2032_至少在两个数组中出现的值
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m := make(map[int]int)
	arr := [][]int{nums1, nums2, nums3}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			value := arr[i][j]
			m[value] = m[value] | (1 << i)
		}
	}
	res := make([]int, 0)
	for k, v := range m {
		if bits.OnesCount(uint(v)) >= 2 {
			res = append(res, k)
		}
	}
	return res
}
