package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{3, 4, 3, 3}))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(nums); i++ {
		sum1 = sum1 + nums[i]
		if _, ok := m[nums[i]]; !ok {
			sum2 = sum2 + nums[i]
		}
		m[nums[i]]++
	}
	return (sum2*3 - sum1) / 2
}
