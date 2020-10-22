package main

import "fmt"

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
}

func findNumbers(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		count := 0
		value := nums[i]
		for value > 0 {
			value = value / 10
			count++
		}
		if count%2 == 0 {
			res++
		}
	}
	return res
}
