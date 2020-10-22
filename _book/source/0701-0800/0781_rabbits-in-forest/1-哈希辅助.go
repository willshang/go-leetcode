package main

import "fmt"

func main() {
	//fmt.Println(numRabbits([]int{1, 1, 2, 2, 0}))
	fmt.Println(numRabbits([]int{1, 1, 0, 0, 0, 1}))
}

// leetcode781_森林中的兔子
func numRabbits(answers []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(answers); i++ {
		value := answers[i]
		if m[value] == 0 {
			res = res + value + 1
		}
		m[value]++
		if m[value] == value+1 {
			m[value] = 0
		}
	}
	return res
}
