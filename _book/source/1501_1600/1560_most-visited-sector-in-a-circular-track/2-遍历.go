package main

import "fmt"

func main() {
	fmt.Println(mostVisited(4, []int{1, 3, 1, 2}))
	fmt.Println(mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
	fmt.Println(mostVisited(7, []int{1, 3, 5, 7}))
}

// leetcode1560_圆形赛道上经过次数最多的扇区
func mostVisited(n int, rounds []int) []int {
	res := make([]int, 0)
	start := rounds[0]
	end := rounds[len(rounds)-1]
	if start <= end {
		for i := start; i <= end; i++ {
			res = append(res, i)
		}
	} else {
		for i := 1; i <= end; i++ {
			res = append(res, i)
		}
		for i := start; i <= n; i++ {
			res = append(res, i)
		}
	}
	return res
}
