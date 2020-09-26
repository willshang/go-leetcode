package main

import (
	"fmt"
)

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}

func average(salary []int) float64 {
	sum := salary[0]
	max := salary[0]
	min := salary[0]
	for i := 1; i < len(salary); i++ {
		sum = sum + salary[i]
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
	}
	return float64(sum-max-min) / float64(len(salary)-2)
}
