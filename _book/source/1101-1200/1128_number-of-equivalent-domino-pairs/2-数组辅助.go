package main

import "fmt"

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}

// leetcode1128_等价多米诺骨牌对的数量
func numEquivDominoPairs(dominoes [][]int) int {
	res := 0
	arr := make([]int, 101)
	for i := 0; i < len(dominoes); i++ {
		a := dominoes[i][0]
		b := dominoes[i][1]
		if a > b {
			a, b = b, a
		}
		res = res + arr[a*10+b]
		arr[a*10+b]++
	}
	return res
}
