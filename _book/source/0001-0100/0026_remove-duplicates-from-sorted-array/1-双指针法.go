package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}

func removeDuplicates(nums []int) int {
	i, j, length := 0, 1, len(nums)
	for ; j < length; j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}
