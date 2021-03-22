package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}

// leetcode1389_按既定顺序创建目标数组
func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(index); i++ {
		for j := len(res) - 1; j > index[i]; j-- {
			res[j] = res[j-1]
		}
		res[index[i]] = nums[i]
	}
	return res
}
