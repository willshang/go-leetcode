package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 3))
}

// leetcode6_Z字形变换
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := []rune(s)
	res := make([]string, numRows)
	flag := -1
	index := 0
	for i := 0; i < len(arr); i++ {
		res[index] = res[index] + string(arr[i])
		if index == 0 || index == numRows-1 {
			flag = -flag
		}
		index = index + flag
	}
	return strings.Join(res, "")
}
