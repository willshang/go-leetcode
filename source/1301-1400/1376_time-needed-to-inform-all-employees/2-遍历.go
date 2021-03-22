package main

import "fmt"

func main() {
	fmt.Println(numOfMinutes(4, 2, []int{3, 3, -1, 2}, []int{0, 0, 162, 914}))
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	res := 0
	for i := 0; i < len(manager); i++ {
		// 没有下属
		if informTime[i] == 0 {
			count := 0
			index := i
			for index != -1 {
				count = count + informTime[index]
				index = manager[index]
			}
			res = max(res, count)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
