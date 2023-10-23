package main

func main() {

}

// leetcode1658_将x减到0的最小操作数
func minOperations(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
	}
	target := sum - x
	left := 0
	right := 0
	res := -1
	cur := 0
	// 和为sum-x的最长连续子数组
	for left < n {
		if right < n {
			cur = cur + nums[right]
			right++
		}
		for cur > target && left < n {
			cur = cur - nums[left]
			left++
		}
		if cur == target {
			res = max(res, right-left)
		}
		if right == n {
			left++
		}
	}
	if res == -1 {
		return -1
	}
	return n - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
