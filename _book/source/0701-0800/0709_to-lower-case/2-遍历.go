package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	fmt.Println(toLowerCase(str))
}

// leetcode709_转换成小写字母
func toLowerCase(str string) string {
	arr := []byte(str)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] - 'A' + 'a'
		}
	}
	return string(arr)
}
