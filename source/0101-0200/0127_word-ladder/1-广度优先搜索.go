package main

import "fmt"

func main() {
	//fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hot", "dog", []string{"hot", "dog"}))
}

// leetcode127_单词接龙
func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]int)
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = 1
	}
	if m[endWord] == 0 {
		return 0
	}
	preMap := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList[i]); j++ {
			newStr := wordList[i][:j] + "*" + wordList[i][j+1:]
			if _, ok := preMap[newStr]; !ok {
				preMap[newStr] = make([]string, 0)
			}
			preMap[newStr] = append(preMap[newStr], wordList[i])
		}
	}
	visited := make(map[string]bool)
	count := 0
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	for len(queue) > 0 {
		count++
		length := len(queue)
		for i := 0; i < length; i++ {
			for j := 0; j < len(beginWord); j++ {
				newStr := queue[i][:j] + "*" + queue[i][j+1:]
				for _, word := range preMap[newStr] {
					if word == endWord {
						return count + 1
					}
					if visited[word] == false {
						visited[word] = true
						queue = append(queue, word)
					}
				}
			}
		}
		queue = queue[length:]
	}
	return 0
}
