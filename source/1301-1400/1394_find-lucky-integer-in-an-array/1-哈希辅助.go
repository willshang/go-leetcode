package main

import "fmt"

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))
}

// leetcode1394_找出数组中的幸运数
func findLucky(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	max := -1
	for i := range m {
		if i == m[i] && max < i {
			max = i
		}
	}
	return max
}
