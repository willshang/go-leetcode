package main

func main() {

}

var maxValue int
var res int

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	maxValue = 0
	res = 0
	for i := 0; i < n; i++ {
		maxValue = maxValue | nums[i]
	}
	dfs(nums, 0, 0)
	return res
}

func dfs(nums []int, index, sum int) {
	if index == len(nums) {
		if sum == maxValue {
			res++
		}
		return
	}
	dfs(nums, index+1, sum)
	dfs(nums, index+1, sum|nums[index])
}
