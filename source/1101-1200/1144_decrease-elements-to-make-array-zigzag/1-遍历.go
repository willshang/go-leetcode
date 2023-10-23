package main

func main() {

}

// leetcode1144_递减元素使数组呈锯齿状
func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 { // 偶数下标小，如果大于左右两边数，需要减去
			left, right := 0, 0
			if i > 0 && nums[i] >= nums[i-1] {
				left = nums[i] - nums[i-1] + 1
			}
			if i < n-1 && nums[i] >= nums[i+1] {
				right = nums[i] - nums[i+1] + 1
			}
			a = a + max(left, right)
		} else { // 奇数下标小，如果大于左右两边数，需要减去
			left, right := 0, 0
			if nums[i] >= nums[i-1] {
				left = nums[i] - nums[i-1] + 1
			}
			if i < n-1 && nums[i] >= nums[i+1] {
				right = nums[i] - nums[i+1] + 1
			}
			b = b + max(left, right)
		}
	}
	return min(a, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
