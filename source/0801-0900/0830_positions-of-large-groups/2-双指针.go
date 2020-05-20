package main

import "fmt"

func main() {
	//fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("xxxxx"))
}

func largeGroupPositions(S string) [][]int {
	res := make([][]int, 0, len(S)/3)
	left, right := 0, 1
	for ; right < len(S); right++ {
		if S[left] != S[right] {
			left = right
			continue
		}
		if right-left+1 == 3 {
			res = append(res, []int{left, right})
		} else if right-left+1 > 3 {
			res[len(res)-1][1] = right
		}
	}
	return res
}
