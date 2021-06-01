package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	origin := []int{1, 1000000, 1, 2, 3, 99, math.MaxInt32}
	res, m := getLSH(origin)
	fmt.Println(origin) // [1 1000000 1 2 3 99 2147483647]
	fmt.Println(res)    // [1 5 1 2 3 4 6]
	fmt.Println(m)      // map[1:1 2:2 3:3 99:4 1000000:5 2147483647:6]
}

func getLSH(a []int) ([]int, map[int]int) {
	n := len(a)
	arr := make([]int, n)
	copy(arr, a)
	sort.Ints(arr) // 排序
	m := make(map[int]int)
	m[arr[0]] = 1
	index := 1
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			index++
			m[arr[i]] = index
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = m[a[i]]
	}
	return res, m
}
