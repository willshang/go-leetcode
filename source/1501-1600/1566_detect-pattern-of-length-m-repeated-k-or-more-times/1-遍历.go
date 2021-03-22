package main

import "fmt"

func main() {
	//fmt.Println(containsPattern([]int{1, 2, 4, 4, 4, 4}, 1, 3))
	fmt.Println(containsPattern([]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2))
	fmt.Println(containsPattern([]int{1, 2, 1, 2, 1, 3}, 2, 3))
	fmt.Println(containsPattern([]int{2, 2}, 1, 2))
}

func containsPattern(arr []int, m int, k int) bool {
	for i := 0; i < len(arr)-m; i++ {
		count := 0
		j := i
		for j+m <= len(arr) {
			flag := true
			for l := 0; l < m; l++ {
				if arr[i+l] != arr[j+l] {
					flag = false
					break
				}
			}
			if flag == true {
				count++
				j = j + m
			} else {
				break
			}
		}
		if count >= k {
			return true
		}
	}
	return false
}
