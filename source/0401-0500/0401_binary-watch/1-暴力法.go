package main

import (
	"fmt"
)

func main() {
	fmt.Println(readBinaryWatch(1))
}

func binCount(num int) int {
	count := make([]int, 0)
	for num != 0 {
		temp := num % 2
		count = append(count, temp)
		num = num / 2
	}
	countNum := 0
	for i := 0; i < len(count); i++ {
		if count[i] == 1 {
			countNum++
		}
	}
	return countNum
}

func readBinaryWatch(num int) []string {
	res := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if binCount(i)+binCount(j) == num {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return res
}
