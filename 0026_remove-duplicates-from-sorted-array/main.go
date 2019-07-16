package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}

/*
//执行用时：124 ms
func removeDuplicates(nums []int) int {
	i , j , length := 0, 1, len(nums)
	for ; j < length; j++{
		if nums[i] == nums[j]{
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i+1
}*/

func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	return len(nums) - count
}
