package main

import "fmt"

func main() {
	fmt.Println(mostVisited(4, []int{1, 3, 1, 2}))
	fmt.Println(mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
	fmt.Println(mostVisited(7, []int{1, 3, 5, 7}))
}

func mostVisited(n int, rounds []int) []int {
	arr := make([]int, n+1)
	res := make([]int, 0)
	max := 0
	arr[rounds[0]]++
	for i := 0; i < len(rounds)-1; i++ {
		start := rounds[i]
		end := rounds[i+1]
		if start < end {
			for j := start + 1; j <= end; j++ {
				arr[j]++
			}
		} else {
			for j := start + 1; j <= n; j++ {
				arr[j]++
			}
			for j := 1; j <= end; j++ {
				arr[j]++
			}
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			res = make([]int, 0)
			res = append(res, i)
		} else if arr[i] == max {
			res = append(res, i)
		}
	}
	return res
}
