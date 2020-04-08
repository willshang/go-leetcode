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
	for i := range res {
		res[i] = 100
	}

	left, right := -n, 2*n
	for i := 0; i < n; i++ {
		j := n - i - 1
		if S[i] == C {
			left = i
		}
		if S[j] == C {
			right = j
		}
		//i从0->n-1 跟左边的C比较得到最近的距离
		//j从n-1->0 跟右边的C比较得到最近的距离
		res[i] = min(res[i], dist(i, left))
		res[j] = min(res[j], dist(j, right))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
