package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}

// leetcode953_验证外星语词典
func isAlienSorted(words []string, order string) bool {
	newWords := make([]string, len(words))
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	for i := 0; i < len(words); i++ {
		str := ""
		for j := 0; j < len(words[i]); j++ {
			str = str + string(m[words[i][j]]+'a')
		}
		newWords[i] = str
	}
	for i := 0; i < len(newWords)-1; i++ {
		if newWords[i] > newWords[i+1] {
			return false
		}
	}
	return true
}
