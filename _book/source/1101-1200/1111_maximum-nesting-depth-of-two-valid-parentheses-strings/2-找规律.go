package main

import "fmt"

func main() {
	fmt.Println(maxDepthAfterSplit("(()())"))
}

func maxDepthAfterSplit(seq string) []int {
	res := make([]int, 0)
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			res = append(res, i%2)
		} else {
			res = append(res, 1-i%2)
		}
	}
	return res
}
