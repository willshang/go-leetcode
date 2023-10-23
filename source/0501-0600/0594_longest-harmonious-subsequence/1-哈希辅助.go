package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(arr))
}

// leetcode594_最长和谐子序列
func findLHS(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n]++
	}
	res := 0
	for key, value := range m {
		value2, ok := m[key+1]
		if ok {
			t := value + value2
			if res < t {
				res = t
			}
		}
	}
	return res
}
