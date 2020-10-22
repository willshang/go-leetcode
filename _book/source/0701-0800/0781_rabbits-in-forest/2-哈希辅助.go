package main

import "fmt"

func main() {
	//fmt.Println(numRabbits([]int{1, 1, 2, 2, 0}))
	fmt.Println(numRabbits([]int{1, 1, 0, 0, 0, 1}))
}

func numRabbits(answers []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(answers); i++ {
		value := answers[i]
		m[value]++
	}
	for k, v := range m {
		target := k + 1
		res = res + v/target*target
		if v%target > 0 {
			res = res + target
		}
	}
	return res
}
