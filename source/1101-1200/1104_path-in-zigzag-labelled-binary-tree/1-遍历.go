package main

import "fmt"

func main() {
	fmt.Println(pathInZigZagTree(14))
}

// leetcode1104_二叉树寻路
func pathInZigZagTree(label int) []int {
	res := make([]int, 0)
	for label > 0 {
		res = append(res, label)
		label = label / 2
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	i := 1
	if len(res)%2 == 0 {
		i = 2
	}
	for ; i < len(res); i = i + 2 {
		res[i] = (1<<i)*3 - 1 - res[i]
	}
	return res
}
