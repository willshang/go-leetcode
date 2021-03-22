package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
}

// leetcode1441_用栈操作构建数组
func buildArray(target []int, n int) []string {
	res := make([]string, 0)
	j := 0
	for i := 1; i <= n; i++ {
		if j >= len(target) {
			break
		}
		if target[j] != i {
			res = append(res, "Push")
			res = append(res, "Pop")
		} else {
			res = append(res, "Push")
			j++
		}
	}
	return res
}
