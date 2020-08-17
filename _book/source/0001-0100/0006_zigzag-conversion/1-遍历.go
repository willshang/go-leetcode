package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := []rune(s)
	total := numRows*2 - 2
	res := make([]string, numRows)
	for i := 0; i < len(arr); i++ {
		index := i % total
		if index < numRows {
			res[index] = res[index] + string(arr[i])
		} else {
			res[total-index] = res[total-index] + string(arr[i])
		}
	}
	return strings.Join(res, "")
}
