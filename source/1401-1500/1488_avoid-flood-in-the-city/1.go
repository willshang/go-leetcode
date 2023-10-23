package main

import "fmt"

func main() {
	//fmt.Println(avoidFlood([]int{1, 2, 3, 4}))
	fmt.Println(avoidFlood([]int{1, 2, 0, 0, 2, 1}))
}

func avoidFlood(rains []int) []int {
	res := make([]int, len(rains))
	queue := make([]int, 0)
	prev := make(map[int]int)
	for i := 0; i < len(rains); i++ {
		// fmt.Println(i, rains[i], queue, prev)
		if rains[i] > 0 {
			if prev[rains[i]] == 0 {
				prev[rains[i]] = rains[i]
				res[i] = -1
			} else {
				if len(queue) > 0 {
					index := queue[0]
					queue = queue[1:]
					res[index] = prev[rains[i]]
					res[i] = -1
					prev[rains[i]] = 0
				} else {
					return nil
				}
			}
		} else {
			queue = append(queue, i)
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			res[i] = 1
		}
	}
	return res
}
