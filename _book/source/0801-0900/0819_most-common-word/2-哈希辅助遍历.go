package main

import (
	"fmt"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	//fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
}

func mostCommonWord(paragraph string, banned []string) string {
	isBanned := make(map[string]bool)
	for _, b := range banned {
		isBanned[b] = true
	}
	count := make(map[string]int)
	length := len(paragraph)
	for i := 0; i < length; i++ {
		for i < length && !isChar(paragraph[i]) {
			i++
		}
		j := i
		temp := ""
		for ; j < length; j++ {
			if !isChar(paragraph[j]) {
				break
			}
			if paragraph[j] >= 'A' && paragraph[j] <= 'Z' {
				temp = temp + string(paragraph[j]-'A'+'a')
			} else {
				temp = temp + string(paragraph[j])
			}
		}
		i = j
		if isBanned[temp] {
			continue
		}
		count[temp]++
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

func isChar(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}
