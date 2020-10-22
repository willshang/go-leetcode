package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getNoZeroIntegers(11))
}

// leetcode1317_将整数转换为两个无零整数的和
func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		if strings.ContainsAny(strconv.Itoa(i), "0") ||
			strings.ContainsAny(strconv.Itoa(n-i), "0") {
			continue
		}
		return []int{i, n - i}
	}
	return nil
}
