package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
}

// leetcode388_文件的最长绝对路径
func lengthLongestPath(input string) int {
	res := 0
	arr := strings.Split(input, "\n")
	temp := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		str := arr[i]
		count := strings.Count(arr[i], "\t")
		str = str[count:]      // 去除\t后的字符串
		if count < len(temp) { // 多个\t的个数小于当前层级，需要去除
			temp = temp[:count]
		}
		if strings.Contains(str, ".") {
			length := len(str) + getLength(temp) + len(temp) // len(temp)个分隔符
			if length > res {
				res = length
			}
		} else {
			temp = append(temp, str)
		}
	}
	return res
}

func getLength(arr []string) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res = res + len(arr[i])
	}
	return res
}
