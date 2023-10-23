package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}

// leetcode1356_根据数字二进制下1的数目排序
func sortByBits(arr []int) []int {
	sort.Ints(arr)
	m := make(map[int][]int, 0)
	for i := 0; i < len(arr); i++ {
		num := countBit(arr[i])
		m[num] = append(m[num], arr[i])
	}
	res := make([]int, 0)
	for i := 0; i < 32; i++ {
		for _, value := range m[i] {
			res = append(res, value)
		}
	}
	return res
}

func countBit(num int) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num = num / 2
	}
	return res
}
