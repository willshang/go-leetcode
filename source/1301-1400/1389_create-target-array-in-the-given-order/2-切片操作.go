package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(index); i++ {
		copy(res[index[i]+1:], res[index[i]:])
		res[index[i]] = nums[i]
	}
	return res
}
