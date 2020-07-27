package main

import "fmt"

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
}

// leetcode970_强整数
func powerfulIntegers(x int, y int, bound int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	if bound < 2 {
		return res
	}
	for i := 1; i < bound; i = i * x {
		for j := 1; i+j <= bound; j = j * y {
			if _, ok := m[i+j]; !ok {
				res = append(res, i+j)
				m[i+j] = 1
			}
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	return res
}
