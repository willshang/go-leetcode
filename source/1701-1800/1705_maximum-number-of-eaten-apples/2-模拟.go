package main

import "fmt"

func main() {
	fmt.Println(eatenApples([]int{3, 0, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 0, 2}))
	//fmt.Println(eatenApples([]int{1}, []int{2}))
}

// leetcode1705_吃苹果的最大数目
func eatenApples(apples []int, days []int) int {
	arr := make([]int, 40000)
	res := 0
	left, right := 0, 0
	for i := 0; ; i++ {
		if i < len(apples) {
			target := i + days[i] // 保质期内
			arr[target] = arr[target] + apples[i]
			if target > right {
				right = target
			}
			if left > target {
				left = target
			}
		} else {
			if left > right {
				break
			}
		}
		// 吃苹果
		for left = i + 1; left <= right; left++ {
			if arr[left] > 0 {
				res++
				arr[left]--
				break
			}
		}
	}
	return res
}
