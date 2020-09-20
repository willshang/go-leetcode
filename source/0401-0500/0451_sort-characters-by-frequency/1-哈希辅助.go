package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("tree"))
}

// leetcode451_根据字符出现频率排序
func frequencySort(s string) string {
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[int(s[i])]++
	}
	arr := make([][2]int, 0)
	for k, v := range m {
		arr = append(arr, [2]int{k, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	res := ""
	for i := range arr {
		for j := 0; j < arr[i][1]; j++ {
			res = res + string(arr[i][0])
		}
	}
	return res
}
