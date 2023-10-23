package main

func main() {

}

// leetcode2321_拼接数组的最大分数
func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return max(solve(nums1, nums2), solve(nums2, nums1)) // 取2种情况最大值
}

func solve(nums1 []int, nums2 []int) int {
	maxDiff := 0
	sum1 := 0
	sum := 0
	// nums1 + maxDiff(num2[left:right]-sum1[left:right])
	// 求maxDiff最大
	// 使用最大子序和来计算maxDiff
	for i := 0; i < len(nums1); i++ {
		sum1 = sum1 + nums1[i]
		// leetcode 53.最大子序和
		sum = max(0, sum+nums2[i]-nums1[i])
		maxDiff = max(maxDiff, sum)
	}
	return sum1 + maxDiff
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
