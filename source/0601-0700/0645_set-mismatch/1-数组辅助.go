package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 2}
	fmt.Println(findErrorNums(nums))
}

func findErrorNums(nums []int) []int {
	newNums := make([]int, len(nums))
	var repeatNum int
	for _, v := range nums {
		if newNums[v-1] != 0 {
			repeatNum = v
		}
		newNums[v-1] = v
	}
	for i, v := range newNums {
		if v == 0 {
			return []int{repeatNum, i + 1}
		}
	}
	return []int{0, 0}
}
