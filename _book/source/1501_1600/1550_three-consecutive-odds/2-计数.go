package main

import "fmt"

func main() {
	fmt.Println(threeConsecutiveOdds([]int{1, 3, 5, 7, 2}))
}

// leetcode1550_存在连续三个奇数的数组
func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			count++
		} else {
			count = 0
		}
		if count == 3 {
			return true
		}
	}
	return false
}
