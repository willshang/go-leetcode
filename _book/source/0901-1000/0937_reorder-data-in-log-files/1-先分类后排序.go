package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))
}

// leetcode937_重新排列日志文件
func reorderLogFiles(logs []string) []string {
	numLogs := make([]string, 0)
	wordLogs := make([]string, 0)
	for key := range logs {
		for i := 0; i < len(logs[key]); i++ {
			if logs[key][i] == ' ' && i != len(logs[key])-1 {
				if strings.ContainsAny(logs[key][i+1:], "0123456789") {
					numLogs = append(numLogs, logs[key])
				} else {
					wordLogs = append(wordLogs, logs[key])
				}
				break
			}
		}
	}
	sort.Slice(wordLogs, func(i, j int) bool {
		firstIndex := strings.Index(wordLogs[i], " ")
		secondIndex := strings.Index(wordLogs[j], " ")
		if wordLogs[i][firstIndex+1:] == wordLogs[j][secondIndex+1:] {
			return wordLogs[i][:firstIndex] < wordLogs[j][:secondIndex]
		}
		return wordLogs[i][firstIndex+1:] < wordLogs[j][secondIndex+1:]
	})
	return append(wordLogs, numLogs...)
}
