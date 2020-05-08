package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(readBinaryWatch(1))
}

// leetcode401_二进制手表
func readBinaryWatch(num int) []string {
	res := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			hour := fmt.Sprintf("%b", i)
			minute := fmt.Sprintf("%b", j)
			if strings.Count(hour, "1")+strings.Count(minute, "1") == num {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return res
}
