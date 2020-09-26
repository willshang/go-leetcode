package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, h int
	var m, w int
	need := make([]int, 0)
	have := make([]int, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		need = append(need, h)
	}
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&w)
		have = append(have, w)
	}
	fmt.Println(deal(need, have))
}

func deal(need, have []int) int {
	sort.Ints(need)
	sort.Ints(have)
	res := 0
	for i, j := 0, 0; i < len(have) && j < len(need); i++ {
		if have[i] >= need[j] {
			res++
			j++
		}
	}
	return res
}
