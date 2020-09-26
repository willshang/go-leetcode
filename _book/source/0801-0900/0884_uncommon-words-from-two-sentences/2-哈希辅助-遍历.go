package main

import (
	"fmt"
)

func main() {
	A := "this apple is sweet"
	B := "this apple is sour"
	fmt.Println(uncommonFromSentences(A, B))
}

func uncommonFromSentences(A string, B string) []string {
	m := map[string]int{}
	A = A + " " + B + " "
	j := 0
	for i := 0; i < len(A); i++ {
		if A[i] == ' ' {
			m[A[j:i]]++
			j = i + 1
		}
	}
	res := make([]string, 0)
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
