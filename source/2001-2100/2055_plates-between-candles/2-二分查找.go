package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sort.SearchInts([]int{1, 3, 9, 20}, 18))
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{
		{2, 5},
		{5, 9},
	}))

	fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{
		{1, 17},
		{4, 5},
	}))
}

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	res := make([]int, len(queries))
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == '|' { // 盘子
			arr = append(arr, i)
		}
	}
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		l := sort.SearchInts(arr, a)
		r := sort.SearchInts(arr, b)
		if r == len(arr) || arr[r] != b { // r超出范围或者找到的下标不是目标数
			r--
		}
		if l < r {
			res[i] = arr[r] - arr[l] - (r - l) // 总个数-盘子个数
		}
	}
	return res
}
