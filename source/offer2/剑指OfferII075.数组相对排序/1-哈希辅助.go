package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	if len(arr2) == 0 {
		sort.Ints(arr1)
		return arr1
	}
	res := make([]int, 0)
	m := make(map[int]int)
	for i := range arr1 {
		m[arr1[i]]++
	}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < m[arr2[i]]; j++ {
			res = append(res, arr2[i])
		}
		m[arr2[i]] = 0
	}
	tempArr := make([]int, 0)
	for key, value := range m {
		for value > 0 {
			tempArr = append(tempArr, key)
			value--
		}
	}
	sort.Ints(tempArr)
	res = append(res, tempArr...)
	return res
}
