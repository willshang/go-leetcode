package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	sum := 0
	singleSum := 0
	for _, v := range nums {
		if m[v] == 0 {
			singleSum = singleSum + v
		}
		m[v] = 1
		sum = sum + v
	}
	return (singleSum*3 - sum) / 2
}
