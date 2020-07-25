package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

func findRepeatedDnaSequences(s string) []string {
	res := make([]string, 0)
	m := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		m[s[i:i+10]]++
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
