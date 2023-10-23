package main

func main() {

}

func buildArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] + 1000*(nums[nums[i]]%1000) // 前3位存储目标值，后3位存储原来值
	}
	for i := 0; i < n; i++ {
		nums[i] = nums[i] / 1000
	}
	return nums
}
