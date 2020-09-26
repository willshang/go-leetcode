package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

// leetcode331_验证二叉树的前序序列化
func isValidSerialization(preorder string) bool {
	arr := strings.Split(preorder, ",")
	slot := 1
	for i := 0; i < len(arr); i++ {
		slot--
		if slot < 0 {
			return false
		}
		if arr[i] != "#" {
			slot = slot + 2
		}
	}
	return slot == 0
}
