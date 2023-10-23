package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"}))
	fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
}

// leetcode819_最常见的单词
func mostCommonWord(paragraph string, banned []string) string {
	isBanned := make(map[string]bool)
	for _, b := range banned {
		isBanned[b] = true
	}
	chars := []string{"!", "?", ",", "'", ";", "."}
	for _, c := range chars {
		paragraph = strings.Replace(paragraph, c, " ", -1)
	}
	p := strings.ToLower(paragraph)
	ss := strings.Fields(p)
	count := make(map[string]int)
	for _, v := range ss {
		if isBanned[v] {
			continue
		}
		count[v]++
	}
	res := ""
	max := 0
	for s, c := range count {
		if max < c {
			max = c
			res = s
		}
	}
	return res
}
