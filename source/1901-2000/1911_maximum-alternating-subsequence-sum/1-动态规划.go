package main

func main() {

}

// leetcode1911_最大子序列交替和
func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	odd, even := 0, nums[0]
	for i := 1; i < n; i++ {
		odd, even = max(even-nums[i], odd), max(odd+nums[i], even)
	}
	return int64(even)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
