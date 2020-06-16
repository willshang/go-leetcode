package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr))
}

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = 1
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[i+1]; !ok {
			res = append(res, i+1)
		}
	}
	return res
}
