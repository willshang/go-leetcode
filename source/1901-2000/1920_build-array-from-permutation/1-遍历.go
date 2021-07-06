package main

func main() {

}

// leetcode1920_基于排列构建数组
func buildArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = nums[nums[i]]
	}
	return res
}
