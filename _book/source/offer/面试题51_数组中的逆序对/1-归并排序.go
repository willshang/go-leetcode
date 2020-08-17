package main

import "fmt"

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}

// 剑指offer_面试题51_数组中的逆序对
var res int

func reversePairs(nums []int) int {
	res = 0
	if len(nums) <= 1 {
		return res
	}
	merge(nums, 0, len(nums)-1)
	return res
}

func merge(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	i, j := left, mid+1
	merge(nums, left, mid)
	merge(nums, mid+1, right)

	temp := make([]int, 0)
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			res = res + mid - i + 1
			temp = append(temp, nums[j])
			j++
		}
	}
	temp = append(temp, nums[i:mid+1]...)
	temp = append(temp, nums[j:right+1]...)
	for key, value := range temp {
		nums[left+key] = value
	}
}
