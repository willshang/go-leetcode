package main

import "fmt"

func main() {
	fmt.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
}

// leetcode1598_文件夹操作日志搜集器
func minOperations(logs []string) int {
	stack := make([]string, 0)
	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if logs[i] != "./" {
			stack = append(stack, logs[i])
		}
	}
	return len(stack)
}
