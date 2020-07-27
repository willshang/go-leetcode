package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "this apple is  sweet"
	B := "this apple is sour"
	fmt.Println(uncommonFromSentences(A, B))
}

// leetcode884_两句话中的不常见单词
func uncommonFromSentences(A string, B string) []string {
	m := map[string]int{}
	arrA := strings.Fields(A)
	arrB := strings.Fields(B)
	for _, v := range arrA {
		m[v]++
	}
	for _, v := range arrB {
		m[v]++
	}
	res := make([]string, 0)
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
