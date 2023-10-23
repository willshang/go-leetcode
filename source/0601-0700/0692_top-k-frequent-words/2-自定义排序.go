package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(arr, 2))
}

// leetcode692_前K个高频单词
func topKFrequent(words []string, k int) []string {
	var res []string
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	var arr []Node
	for k, v := range m {
		arr = append(arr, Node{
			str: k,
			num: v,
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].num == arr[j].num {
			return arr[i].str < arr[j].str
		}
		return arr[i].num > arr[j].num
	})
	for i := 0; i < k; i++ {
		res = append(res, arr[i].str)
	}
	return res
}

type Node struct {
	str string
	num int
}
