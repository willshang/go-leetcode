package main

func main() {

}

// leetcode1671_得到山形数组的最少删除次数
func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	res := 0
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				left[i] = max(left[j]+1, left[i])
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		right[i] = 0
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] {
				right[i] = max(right[j]+1, right[i])
			}
		}
	}
	for i := 1; i < n-1; i++ {
		res = max(res, left[i]+right[i]+1)
	}
	return n - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
