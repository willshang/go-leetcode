package main

func main() {

}

// leetcode1855_下标对中的最大距离
func maxDistance(nums1 []int, nums2 []int) int {
	res := 0
	j := 0
	for i := 0; i < len(nums1); i++ {
		for j < len(nums2) && nums1[i] <= nums2[j] {
			j++
		}
		if j-i-1 > res {
			res = j - i - 1
		}
	}
	return res
}
