package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = true
	}
	if m[endWord] == false {
		return 0
	}
	list := make([]string, 0)
	list = append(list, beginWord)
	res := 0
	for len(list) > 0 {
		res++
		length := len(list)
		for i := length - 1; i >= 0; i-- {

		}
		for i := 0; i < len(beginWord); i++ {
			char := beginWord[i]
			for j := 0; j < 26; j++ {
				if j+'a' == int(char) {
					continue
				}
			}
		}
	}
	return res
}
