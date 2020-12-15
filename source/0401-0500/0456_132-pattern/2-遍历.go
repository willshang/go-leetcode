package main

func main() {

}

// leetcode456_132模式
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	minArr := make([]int, len(nums)) // minArr[i]是[0,i]中的最小值
	minArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		minArr[i] = min(nums[i], minArr[i-1])
	}
	for j := len(nums) - 2; j >= 0; j-- {
		if nums[j] != minArr[j] {
			for k := j + 1; k < len(nums); k++ {
				if nums[k] > minArr[j] && nums[k] < nums[j] {
					return true
				}
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
