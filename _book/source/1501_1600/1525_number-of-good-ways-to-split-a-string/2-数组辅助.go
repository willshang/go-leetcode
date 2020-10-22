package main

import "fmt"

func main() {
	fmt.Println(numSplits("abcd"))
}

func numSplits(s string) int {
	left := [256]int{}
	right := [256]int{}
	leftCount := 0
	rightCount := 0
	for i := 0; i < len(s); i++ {
		left[s[i]]++
		if left[s[i]] == 1 {
			leftCount++
		}
	}
	res := 0
	for i := 0; i < len(s); i++ {
		left[s[i]]--
		right[s[i]]++
		if left[s[i]] == 0 {
			leftCount--
		}
		if right[s[i]] == 1 {
			rightCount++
		}
		if leftCount == rightCount {
			res++
		}
	}
	return res
}
