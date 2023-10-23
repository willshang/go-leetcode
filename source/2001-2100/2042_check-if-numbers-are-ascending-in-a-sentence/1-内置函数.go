package main

import (
	"strconv"
	"strings"
)

func main() {

}

// leetcode2042_检查句子中的数字是否递增
func areNumbersAscending(s string) bool {
	arr := strings.Split(s, " ")
	prev := -1
	for i := 0; i < len(arr); i++ {
		if '0' <= arr[i][0] && arr[i][0] <= '9' {
			value, _ := strconv.Atoi(arr[i])
			if value > prev {
				prev = value
			} else {
				return false
			}
		}
	}
	return true
}
