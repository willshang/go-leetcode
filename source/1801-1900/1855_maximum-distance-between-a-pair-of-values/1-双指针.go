package main

func main() {

}

func maxDistance(nums1 []int, nums2 []int) int {
	res := 0
	i := 0
	for j := 0; j < len(nums2); j++ {
		for i < len(nums1) && nums2[j] < nums1[i] {
			i++
		}
		if i < len(nums1) {
			if j-i > res {
				res = j - i
			}
		}
	}
	return res
}
