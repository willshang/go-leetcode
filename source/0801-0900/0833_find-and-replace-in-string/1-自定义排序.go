package main

import (
	"fmt"
	"sort"
)

func main() {
	a := "jjievdtjfb"
	arr := []int{4, 6, 1}
	a1 := []string{"md", "tjgb", "jf"}
	a2 := []string{"foe", "oov", "e"}
	fmt.Println(findReplaceString(a, arr, a1, a2))
}

type Node struct {
	index  int
	source string
	target string
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	arr := make([]Node, 0)
	for i := 0; i < len(indexes); i++ {
		arr = append(arr, Node{
			index:  indexes[i],
			source: sources[i],
			target: targets[i],
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].index < arr[j].index
	})
	res := ""
	left := 0
	for i := 0; i < len(arr); i++ {
		if left < arr[i].index {
			res = res + S[left:arr[i].index]
			left = arr[i].index
		}
		start := arr[i].index
		end := arr[i].index + len(arr[i].source)
		if end <= len(S) && S[start:end] == arr[i].source {
			res = res + arr[i].target
			left = end
		}
	}
	if left < len(S) {
		res = res + S[left:]
	}
	return res
}
