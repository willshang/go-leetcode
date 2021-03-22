package main

import "fmt"

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}

func numIdenticalPairs(nums []int) int {
	res := 0
	m := make(map[int]int, 101)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, v := range m {
		if v > 0 {
			res = res + v*(v-1)/2
		}
	}
	return res
}
