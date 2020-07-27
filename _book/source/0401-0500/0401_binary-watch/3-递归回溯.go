package main

import (
	"fmt"
)

func main() {
	fmt.Println(readBinaryWatch(3))
}

func readBinaryWatch(num int) []string {
	res := make([]string, 0)
	ledS := make([]bool, 10)

	var dfs func(int, int)
	dfs = func(idx, num int) {
		if num == 0 {
			// 满足条件
			m, h := get(ledS[:6]), get(ledS[6:])
			if h < 12 && m < 60 {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
			return
		}
		for i := idx; i < 11-num; i++ {
			ledS[i] = true
			dfs(i+1, num-1)
			ledS[i] = false
		}
	}
	dfs(0, num)
	return res
}

func get(ledS []bool) int {
	bs := []int{1, 2, 4, 8, 16, 32}
	var sum int
	size := len(ledS)
	for i := 0; i < size; i++ {
		if ledS[i] {
			sum += bs[i]
		}
	}
	return sum
}
