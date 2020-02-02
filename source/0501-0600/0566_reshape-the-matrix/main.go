package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrixReshape(arr, 4, 1))
}
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}

	if len(nums[0]) == 0 {
		return nums
	}

	if len(nums)*(len(nums[0])) != r*c {
		return nums
	}
	if len(nums) == r && len(nums[0]) == c {
		return nums
	}

	res := make([][]int, r)
	count, col := 0, len(nums[0])
	for i := range res {
		res[i] = make([]int, c)

		for j := range res[i] {
			res[i][j] = nums[count/col][count%col]
			count++
		}
	}
	return res
}
