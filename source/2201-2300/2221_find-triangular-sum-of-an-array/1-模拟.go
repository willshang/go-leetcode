package main

func main() {

}

// leetcode2221_数组的三角和
func triangularSum(nums []int) int {
	for len(nums) > 1 {
		arr := make([]int, len(nums)-1)
		for i := 0; i < len(nums)-1; i++ {
			arr[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = arr
	}
	return nums[0]
}
