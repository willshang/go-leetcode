package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumOperations([]int{1, 1, 1, 1, 1}))
}

func minimumOperations(nums []int) int {
	m := [2]map[int]int{}
	for i := 0; i < 2; i++ {
		m[i] = make(map[int]int)
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			m[0][nums[i]]++
		} else {
			m[1][nums[i]]++
		}
	}
	a := make([][2]int, 0)
	b := make([][2]int, 0)
	for i := 0; i < 2; i++ {
		temp := make([][2]int, 2) // 0补足
		for k, v := range m[i] {
			temp = append(temp, [2]int{k, v})
		}
		sort.Slice(temp, func(i, j int) bool {
			return temp[i][1] > temp[j][1]
		})
		if i%2 == 0 {
			a = temp[:2]
		} else {
			b = temp[:2]
		}
	}
	if a[0][0] == b[0][0] {
		return len(nums) - max(a[0][1]+b[1][1], a[1][1]+b[0][1])
	}
	return len(nums) - a[0][1] - b[0][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
