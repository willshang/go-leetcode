package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	left := 1 << n
	right := 1 << (n + 1)
	for i := left; i < right; i++ {
		temp := make([]int, 0)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				temp = append(temp, nums[j])
			}
		}
		res = append(res, temp)
	}
	return res
}
