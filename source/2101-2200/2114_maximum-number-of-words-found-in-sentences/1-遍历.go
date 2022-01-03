package main

import "strings"

func main() {

}

// leetcode2114_句子中的最多单词数
func mostWordsFound(sentences []string) int {
	res := 0
	for i := 0; i < len(sentences); i++ {
		arr := strings.Fields(sentences[i])
		if len(arr) > res {
			res = len(arr)
		}
	}
	return res
}
