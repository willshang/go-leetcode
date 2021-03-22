package main

import "fmt"

func main() {
	fmt.Println(countLargestGroup(13))
}

// leetcode1399_统计最大组的数目
func countLargestGroup(n int) int {
	if n < 10 {
		return n
	}
	m := make(map[int]int, 50)
	max := 0
	for i := 1; i <= n; i++ {
		value := sum(i)
		m[value]++
		if m[value] > max {
			max = m[value]
		}
	}
	res := 0
	for i := range m {
		if m[i] == max {
			res++
		}
	}
	return res
}

func sum(n int) int {
	res := 0
	for n > 0 {
		res = res + n%10
		n = n / 10
	}
	return res
}
