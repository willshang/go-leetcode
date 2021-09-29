package main

func main() {

}

// leetcode2016_增量元素之间的最大差值
func maximumDifference(nums []int) int {
	res := 0
	minValue := nums[0]
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i]-minValue)
		minValue = min(minValue, nums[i])
	}
	if res == 0 {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
