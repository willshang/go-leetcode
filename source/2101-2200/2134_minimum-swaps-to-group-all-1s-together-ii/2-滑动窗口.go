package main

func main() {

}

// leetcode2134_最少交换次数来组合所有的1II
func minSwaps(nums []int) int {
	n := len(nums)
	count := 0 // 1的个数=>滑动窗口的大小
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count++
		}
	}
	res := n // 枚举所有起点，统计滑动窗口内0的个数
	nums = append(nums, nums...)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + (1 - nums[i])
		if i >= count {
			sum = sum - (1 - nums[i-count])
			res = min(res, sum)
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
