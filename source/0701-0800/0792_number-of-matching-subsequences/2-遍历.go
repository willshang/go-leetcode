package main

import "fmt"

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"ace"}))
}

func numMatchingSubseq(S string, words []string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		if len(words[i]) > len(S) {
			continue
		}
		k := 0
		for j := 0; j < len(S); j++ {
			if S[j] == words[i][k] {
				k++
				if k == len(words[i]) {
					res++
					break
				}
			}
		}
	}
	return res
}
