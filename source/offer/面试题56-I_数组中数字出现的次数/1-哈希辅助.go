package main

import "fmt"

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}

func singleNumbers(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	res := make([]int, 0)
	for i := range m {
		if m[i] == 1 {
			res = append(res, i)
		}
	}
	return res
}
