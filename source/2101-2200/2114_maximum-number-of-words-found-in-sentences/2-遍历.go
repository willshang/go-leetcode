package main

import "strings"

func main() {

}

func mostWordsFound(sentences []string) int {
	res := 0
	for i := 0; i < len(sentences); i++ {
		count := strings.Count(sentences[i], " ") + 1
		if count > res {
			res = count
		}
	}
	return res
}
