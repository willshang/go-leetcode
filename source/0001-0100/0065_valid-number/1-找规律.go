package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber("0"))
}

// leetcode65_有效数字
func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if s == "" || len(s) == 0 || len(s) == 0 {
		return false
	}
	arr := []byte(s)
	i := 0
	numeric := scanInteger(&arr, &i)
	if i < len(arr) && arr[i] == '.' {
		i++
		numeric = scanUnsignedInteger(&arr, &i) || numeric
	}
	if i < len(arr) && (arr[i] == 'e' || arr[i] == 'E') {
		i++
		numeric = numeric && scanInteger(&arr, &i)
	}
	return numeric && len(arr) == i
}

func scanInteger(arr *[]byte, index *int) bool {
	if len(*arr) <= *index {
		return false
	}
	if (*arr)[*index] == '+' || (*arr)[*index] == '-' {
		*index++
	}
	return scanUnsignedInteger(arr, index)
}

func scanUnsignedInteger(arr *[]byte, index *int) bool {
	j := *index
	for *index < len(*arr) {
		if (*arr)[*index] < '0' || (*arr)[*index] > '9' {
			break
		}
		*index++
	}
	return j < *index
}
