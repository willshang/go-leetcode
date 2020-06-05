package main

import "fmt"

func main() {
	S := "loveleetcode"
	C := 'e'
	fmt.Println(shortestToChar(S, byte(C)))
}

func shortestToChar(S string, C byte) []int {
	n := len(S)
	res := make([]int, n)
	arr := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			arr = append(arr, i)
		}
	}
	for i := 0; i < n; i++ {
		min := n
		for _, value := range arr {
			if value == i {
				min = 0
				break
			}
			if min > dist(i, value) {
				min = dist(i, value)
			}
		}
		res[i] = min
	}
	return res
}

func dist(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
