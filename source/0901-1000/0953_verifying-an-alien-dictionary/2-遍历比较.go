package main

import "fmt"

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}

func isAlienSorted(words []string, order string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		length := len(words[i])
		if len(words[i+1]) < length {
			length = len(words[i+1])
		}
		for j := 0; j < length; j++ {
			if m[words[i][j]] < m[words[i+1][j]] {
				break
			}
			if m[words[i][j]] > m[words[i+1][j]] {
				return false
			}
			if j == length-1 {
				if len(words[i]) > len(words[i+1]) {
					return false
				}
			}
		}
	}
	return true
}
