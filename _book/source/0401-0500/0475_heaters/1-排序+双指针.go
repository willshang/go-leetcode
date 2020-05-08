package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(findRadius([]int{4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8}))
	fmt.Println(findRadius([]int{4, 5, 6, 7, 8, 9, 10}, []int{0, 1, 14, 19}))
}

// leetcode475_供暖器
func findRadius(houses []int, heaters []int) int {
	if len(heaters) == 0 {
		return 0
	}
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	j := 0
	for i := 0; i < len(houses); i++ {
		// 找到最近的一个供暖器, >=确保出现重复的供暖器会往后走
		for j < len(heaters)-1 &&
			Abs(houses[i], heaters[j]) >= Abs(houses[i], heaters[j+1]) {
			j++
		}
		res = Max(Abs(houses[i], heaters[j]), res)
	}
	return res
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
