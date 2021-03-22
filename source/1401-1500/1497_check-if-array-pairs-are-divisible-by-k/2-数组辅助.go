package main

import "fmt"

func main() {
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
}

// leetcode1497_检查数组对是否可以被k整除
func canArrange(arr []int, k int) bool {
	temp := make([]int, k)
	for i := 0; i < len(arr); i++ {
		value := ((arr[i] % k) + k) % k
		temp[value]++
	}
	for i := 1; i < k; i++ {
		if temp[i] != temp[k-i] {
			return false
		}
	}
	return temp[0]%2 == 0
}
