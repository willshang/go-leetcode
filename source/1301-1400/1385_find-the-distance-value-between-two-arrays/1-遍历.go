package main

import "fmt"

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}

// leetcode1385_两个数组间的距离值
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := 0; i < len(arr1); i++ {
		flag := true
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i], arr2[j]) <= d {
				flag = false
			}
		}
		if flag == true {
			res++
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
