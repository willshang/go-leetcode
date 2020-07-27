package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
}

// leetcode71_简化路径
func simplifyPath(path string) string {
	stack := make([]string, 0)
	arr := strings.Split(path, "/")
	for i := 0; i < len(arr); i++ {
		if arr[i] == "." || arr[i] == "" {
			continue
		}
		if arr[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, arr[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}
