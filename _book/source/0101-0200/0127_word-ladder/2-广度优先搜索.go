package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]int)
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = 1
	}
	if m[endWord] == 0 {
		return 0
	}
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	count := 0
	for len(queue) > 0 {
		count++
		length := len(queue)
		for i := 0; i < length; i++ {
			for _, word := range wordList {
				diff := 0
				for j := 0; j < len(queue[i]); j++ {
					if queue[i][j] != word[j] {
						diff++
					}
					if diff > 1 {
						break
					}
				}
				if diff == 1 && m[word] != 2 {
					if word == endWord {
						return count + 1
					}
					m[word] = 2
					queue = append(queue, word)
				}
			}
		}
		queue = queue[length:]
	}
	return 0
}
