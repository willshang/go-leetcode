package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr))
}

/*
# 遍历过程
init = [4 3 2 7 8 2 3 1]
0 4 [4 3 2 -7 8 2 3 1]
1 3 [4 3 -2 -7 8 2 3 1]
2 2 [4 -3 -2 -7 8 2 3 1]
3 7 [4 -3 -2 -7 8 2 -3 1]
4 8 [4 -3 -2 -7 8 2 -3 -1]
5 2 [4 -3 -2 -7 8 2 -3 -1]
6 3 [4 -3 -2 -7 8 2 -3 -1]
7 1 [-4 -3 -2 -7 8 2 -3 -1]
*/
// leetcode448_找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		if value < 0 {
			value = -value
		}
		if nums[value-1] > 0 {
			nums[value-1] = -nums[value-1]
		}
	}
	res := make([]int, 0)
	for key, value := range nums {
		if value > 0 {
			res = append(res, key+1)
		}
	}
	return res
}
