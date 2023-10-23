package main

func main() {

}

// leetcode1800_最大升序子数组和
func maxAscendingSum(nums []int) int {
	res, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}
		res = max(res, sum)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
