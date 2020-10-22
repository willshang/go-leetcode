package main

import "fmt"

func main() {
	//fmt.Println(missingTwo([]int{2, 3, 4}))
	fmt.Println(missingTwo([]int{3}))
}

func missingTwo(nums []int) []int {
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = temp ^ nums[i]
	}
	for i := 1; i <= len(nums)+2; i++ {
		temp = temp ^ i
	}
	a := 0
	diff := temp & (-temp)
	for i := 1; i <= len(nums)+2; i++ {
		if diff&i != 0 {
			a = a ^ i
		}
	}
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] != 0 {
			a = a ^ nums[i]
		}
	}
	return []int{a, a ^ temp}
}
