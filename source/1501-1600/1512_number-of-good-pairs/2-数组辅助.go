package main

import "fmt"

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}

func numIdenticalPairs(nums []int) int {
	res := 0
	arr := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		res = res + arr[nums[i]]
		arr[nums[i]]++
	}
	return res
}
