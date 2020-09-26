package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

func isValidSerialization(preorder string) bool {
	arr := strings.Split(preorder, ",")
	stack := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] == "#" && arr[i] == "#" {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return len(stack) == 1 && stack[0] == "#"
}
