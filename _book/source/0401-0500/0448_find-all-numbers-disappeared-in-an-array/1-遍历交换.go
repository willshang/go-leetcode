package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr))
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := make([]int, 0)
	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
