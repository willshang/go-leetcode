package main

import "fmt"

func main() {
	fmt.Println(minDays(10))
}

func minDays(n int) int {
	m := make(map[int]bool)
	queue := make([]int, 0)
	queue = append(queue, n)
	res := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			day := queue[i]
			if day == 0 {
				return res
			}
			if day%3 == 0 && m[day/3] == false {
				queue = append(queue, day/3)
				m[day/3] = true
			}
			if day%2 == 0 && m[day/2] == false {
				queue = append(queue, day/2)
				m[day/2] = true
			}
			if m[day-1] == false {
				queue = append(queue, day-1)
				m[day-1] = true
			}
		}
		res++
		queue = queue[length:]
	}
	return res
}
