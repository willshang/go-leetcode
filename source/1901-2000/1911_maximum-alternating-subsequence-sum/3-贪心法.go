package main

func main() {

}

func maxAlternatingSum(nums []int) int64 {
	// leetcode122.买卖股票的最佳时机II
	nums = append([]int{0}, nums...) // 补零
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			res = res + nums[i] - nums[i-1]
		}
	}
	return int64(res)
}
