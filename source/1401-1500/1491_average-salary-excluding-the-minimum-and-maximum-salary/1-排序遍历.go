package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}

// leetcode1491_去掉最低工资和最高工资后的工资平均值
func average(salary []int) float64 {
	sort.Ints(salary)
	sum := 0
	for i := 1; i < len(salary)-1; i++ {
		sum = sum + salary[i]
	}
	return float64(sum) / float64(len(salary)-2)
}
