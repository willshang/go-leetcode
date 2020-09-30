package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
}

// leetcode1481_不同整数的最少数目
func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	temp := make([]int, 0)
	for _, v := range m {
		temp = append(temp, v)
	}
	sort.Ints(temp)
	res := len(temp)
	for i := 0; i < len(temp); i++ {
		if k >= temp[i] {
			res--
			k = k - temp[i]
		} else {
			break
		}
	}
	return res
}
