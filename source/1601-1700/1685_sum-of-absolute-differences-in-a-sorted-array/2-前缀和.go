package main

func main() {

}

// leetcode1685_有序数组中差绝对值之和
func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	arr := make([]int, 0)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
		arr = append(arr, sum)
	}
	res = append(res, sum-n*nums[0])
	for i := 1; i < n; i++ {
		left := nums[i]*i - arr[i-1]              // 左边和
		right := (sum - arr[i]) - nums[i]*(n-1-i) // 右边和
		res = append(res, left+right)
	}
	return res
}
