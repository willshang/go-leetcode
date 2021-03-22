package main

import "fmt"

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
}

// leetcode1450_在既定时间做作业的学生人数
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0
	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}
