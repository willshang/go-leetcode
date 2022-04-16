package main

func main() {

}

func triangularSum(nums []int) int {
	for n := len(nums) - 1; n > 0; n-- {
		for i := 0; i < n; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
	}
	return nums[0]
}
