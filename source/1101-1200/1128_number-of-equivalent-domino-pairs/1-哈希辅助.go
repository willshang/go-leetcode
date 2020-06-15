package main

import "fmt"

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}

func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[string]int)
	for i := 0; i < len(dominoes); i++ {
		a := dominoes[i][0]
		b := dominoes[i][1]
		if a > b {
			a, b = b, a
		}
		m[fmt.Sprintf("%d,%d", a, b)]++
	}
	res := 0
	for _, v := range m {
		res = res + v*(v-1)/2
	}
	return res
}
