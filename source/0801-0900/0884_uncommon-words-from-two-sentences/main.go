package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "this apple is sweet"
	B := "this apple is sour"
	fmt.Println(uncommonFromSentences(A, B))
}
func uncommonFromSentences(A string, B string) []string {
	m := map[string]int{}
	arrA := strings.Split(A, " ")
	arrB := strings.Split(B, " ")

	for _, v := range arrA {
		m[v]++
	}
	for _, v := range arrB {
		m[v]++
	}

	res := []string{}

	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
