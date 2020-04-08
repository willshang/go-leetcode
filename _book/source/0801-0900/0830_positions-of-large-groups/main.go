package main

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
}
func largeGroupPositions(S string) [][]int {
	res := make([][]int, 0, len(S)/3)
	l, r := 0, 1

	for ; r < len(S); r++ {
		if S[l] != S[r] {
			l = r
			continue
		}
		if r-l+1 == 3 {
			res = append(res, []int{l, r})
		} else if r-l+1 > 3 {
			res[len(res)-1][1] = r
		}
	}
	return res
}
