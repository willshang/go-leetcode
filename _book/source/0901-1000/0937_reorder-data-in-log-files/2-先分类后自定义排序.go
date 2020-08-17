package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))
}

type Logs []string

func (l Logs) Len() int {
	return len(l)
}

func (l Logs) Less(i, j int) bool {
	firstIndex := strings.Index(l[i], " ")
	secondIndex := strings.Index(l[j], " ")
	if l[i][firstIndex+1:] == l[j][secondIndex+1:] {
		return l[i][:firstIndex] < l[j][:secondIndex]
	}
	return l[i][firstIndex+1:] < l[j][secondIndex+1:]
}

func (l Logs) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

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
	sort.Sort(Logs(wordLogs))
	return append(wordLogs, numLogs...)
}
