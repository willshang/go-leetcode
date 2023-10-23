package main

import (
	"fmt"
)

func main() {
	//fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string]int)
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = 1
	}
	if m[endWord] == 0 {
		return res
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
	pathQueue := make([][]string, 0)
	pathQueue = append(pathQueue, []string{beginWord})
	flag := false
	for len(queue) > 0 {
		count++
		length := len(queue)
		fmt.Println(pathQueue, queue)
		fmt.Println()
		for i := 0; i < length; i++ {
			fmt.Println("当前", queue[i])
			for j := 0; j < len(beginWord); j++ {
				newStr := queue[i][:j] + "*" + queue[i][j+1:]
				for _, word := range preMap[newStr] {
					if word == endWord {
						fmt.Println(11, newStr, word)
					}
					if word == endWord {
						fmt.Println(pathQueue[i])
						flag = true
						if flag == true {
							res = append(res, append(pathQueue[i], endWord))
						}
					}
					if visited[word] == false {
						visited[word] = true
						newArr := append(pathQueue[i], word)
						pathQueue = append(pathQueue, newArr)
						queue = append(queue, word)
					}
				}
			}
		}
		if flag == true {
			return res
		}
		queue = queue[length:]
		pathQueue = pathQueue[length:]
	}
	return res
}
