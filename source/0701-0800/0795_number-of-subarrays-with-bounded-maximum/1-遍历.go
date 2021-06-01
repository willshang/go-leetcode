package main

func main() {

}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	// L~R范围的组合数=0~R范围的组合数- 0~L-1范围的组合数
	return count(nums, right) - count(nums, left-1)
}

func count(nums []int, target int) int {
	res := 0
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= target {
			total++
		} else {
			total = 0
		}
		res = res + total
	}
	return res
}
