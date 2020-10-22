package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, h int
	var m, w int
	hArr := make([]int, 0)
	wArr := make([]int, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		hArr = append(hArr, h)
	}
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&w)
		wArr = append(wArr, w)
	}
	fmt.Println(deal(hArr, wArr))
}

func deal(hArr, wArr []int) int {
	sort.Ints(hArr)
	sort.Ints(wArr)
	res := 0
	index := len(hArr) - 1
	for i := len(wArr) - 1; i >= 0; i-- {
		for j := index; j >= 0; j-- {
			if wArr[i] >= hArr[j] {
				res++
				index = j - 1
				break
			}
		}
	}
	return res
}
