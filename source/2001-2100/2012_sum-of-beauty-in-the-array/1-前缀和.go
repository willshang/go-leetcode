package main

func main() {

}

// leetcode2012_数组美丽值求和
func sumOfBeauties(nums []int) int {
	res := 0
	n := len(nums)
	arrA := make([]int, n)
	arrA[0] = nums[0]
	for i := 1; i < n; i++ {
		arrA[i] = max(nums[i], arrA[i-1])
	}
	arrB := make([]int, n)
	arrB[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		arrB[i] = min(nums[i], arrB[i+1])
	}
	for i := 1; i <= n-2; i++ {
		if arrA[i-1] < nums[i] && nums[i] < arrB[i+1] {
			res = res + 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			res = res + 1
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
