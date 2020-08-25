package main

import (
	"fmt"
)

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}

func splitArray(nums []int, m int) int {
	left, right := 0, 0 // 最小值，最大值
	for i := 0; i < len(nums); i++ {
		right = right + nums[i]
		if left < nums[i] {
			left = nums[i]
		}
	}
	for left < right {
		mid := left + (right-left)/2
		// 分割数太多，继续尝试
		if check(nums, mid, m) > m {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 区间和的最大值为target时，所得出的区间数
func check(arr []int, target int, m int) int {
	sum := 0
	count := 1
	for i := 0; i < len(arr); i++ {
		// 当sum加上当前值超过了target
		// 我们就把当前取的值作为新的一段分割子数组的开头，并将count加1
		if sum+arr[i] > target {
			count++
			sum = 0
		}
		sum = sum + arr[i]
	}
	return count
}
