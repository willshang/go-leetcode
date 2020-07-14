package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		tempArr := make([]int, len(nums)-1)
		copy(tempArr[0:], nums[:i])
		copy(tempArr[i:], nums[i+1:])
		arr := permute(tempArr)
		for _, v := range arr {
			res = append(res, append(v, nums[i]))
		}
	}
	return res
}
