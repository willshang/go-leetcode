package main

func main() {

}

// leetcode1863_找出所有子集的异或总和再求和
var res int

func subsetXORSum(nums []int) int {
	res = 0
	dfs(nums, 0, 0)
	return res
}

func dfs(nums []int, sum int, index int) {
	if index >= len(nums) {
		res = res + sum
		return
	}
	dfs(nums, sum, index+1)
	dfs(nums, sum^nums[index], index+1)
}
