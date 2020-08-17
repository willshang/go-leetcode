package main

import "fmt"

func main() {
	fmt.Println(divingBoard(1, 2, 3))
}

// 程序员面试金典16.11_跳水板
func divingBoard(shorter int, longer int, k int) []int {
	res := make([]int, 0)
	if k == 0 {
		return res
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	for i := 0; i <= k; i++ {
		res = append(res, shorter*(k-i)+longer*i)
	}
	return res
}
