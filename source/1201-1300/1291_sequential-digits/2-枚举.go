package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(sequentialDigits(100, 300))
}

// leetcode1291_顺次数
func sequentialDigits(low int, high int) []int {
	res := make([]int, 0)
	str := "123456789"
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			num, _ := strconv.Atoi(str[i:j])
			if num >= low && num <= high {
				res = append(res, num)
			}
		}
	}
	sort.Ints(res)
	return res
}
