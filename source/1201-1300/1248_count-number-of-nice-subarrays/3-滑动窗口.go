package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
}

func numberOfSubarrays(nums []int, k int) int {
	res := 0
	left, right := 0, 0
	count := 0
	for right < len(nums) {
		if nums[right]%2 == 1 {
			count++
		}
		right++
		if count == k {
			temp := right
			for right < len(nums) && nums[right]%2 == 0 {
				right++
			}
			totalRight := right - temp + 1
			totalLeft := 1
			for nums[left]%2 == 0 {
				left++
				totalLeft++
			}
			res = res + totalLeft*totalRight
			count--
			left++
		}
	}
	return res
}
