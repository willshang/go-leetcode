package main

func main() {

}

func nextPermutation(nums []int) {
	n := len(nums)
	left := n - 2
	// 以12385764为例，从后往前找到5<7 的升序情况，目标值为左边的数5
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left >= 0 { // 存在升序的情况
		right := n - 1
		// 从后往前，找到第一个大于目标值的数，如6>5，然后交换
		for right >= 0 && nums[right] <= nums[left] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	reverse(nums, left+1, n-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
