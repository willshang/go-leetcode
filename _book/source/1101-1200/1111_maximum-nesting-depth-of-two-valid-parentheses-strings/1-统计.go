package main

import "fmt"

func main() {
	fmt.Println(maxDepthAfterSplit("(()())"))
}

func maxDepthAfterSplit(seq string) []int {
	res := make([]int, 0)
	level := 0
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			level++
			res = append(res, level%2)
		} else {
			res = append(res, level%2)
			level--
		}
	}
	return res
}
