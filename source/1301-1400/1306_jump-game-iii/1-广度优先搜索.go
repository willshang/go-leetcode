package main

import "fmt"

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
}

func canReach(arr []int, start int) bool {
	m := make(map[int]bool)
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		length := len(queue)
		for j := 0; j < length; j++ {
			i := queue[j]
			if m[i] == false {
				m[i] = true
				if i+arr[i] < len(arr) {
					if arr[i+arr[i]] == 0 {
						return true
					}
					queue = append(queue, i+arr[i])
				}
				if i-arr[i] >= 0 {
					if arr[i-arr[i]] == 0 {
						return true
					}
					queue = append(queue, i-arr[i])
				}
			}
		}
		queue = queue[length:]
	}
	return false
}
