package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B'}, 2))
}

func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		arr[tasks[i]-'A']++
	}
	sort.Ints(arr)
	res := 0
	for arr[25] > 0 {
		i := 0
		for i <= n { // 每次安排n+1个
			if arr[25] == 0 {
				break
			}
			if i < 26 && arr[25-i] > 0 {
				arr[25-i]--
			}
			res++
			i++
		}
		sort.Ints(arr)
	}
	return res
}
