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
	for i := 0; i < n; i++ {
		if S[i] == C {
			res[i] = 0
			continue
		}
		min := n
		for j := i + 1; j < n; j++ {
			if S[j] == C {
				if min > j-i {
					min = j - i
				}
				break
			}
		}
		for j := i - 1; j >= 0; j-- {
			if S[j] == C {
				if min > i-j {
					min = i - j
				}
				break
			}
		}
		res[i] = min
	}
	return res
}
