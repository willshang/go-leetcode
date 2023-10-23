package main

func main() {

}

// leetcode1929_数组串联
func getConcatenation(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}
