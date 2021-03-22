package main

func main() {

}

// leetcode1537_最大得分
func maxSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	a, b := 0, 0
	i, j := 0, 0
	for i < n || j < m {
		if i < n && j < m {
			if nums1[i] < nums2[j] {
				a = a + nums1[i]
				i++
			} else if nums1[i] > nums2[j] {
				b = b + nums2[j]
				j++
			} else {
				temp := max(a, b) + nums1[i] // 遇到相同值，取较大值
				a = temp
				b = temp
				i++
				j++
			}
		} else if i < n {
			a = a + nums1[i]
			i++
		} else if j < m {
			b = b + nums2[j]
			j++
		}
	}
	return max(a, b) % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
