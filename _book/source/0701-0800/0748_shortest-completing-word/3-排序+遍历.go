package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	// fmt.Println(shortestCompletingWord("1s3 456", []string{"looks","pest","stew","show"}))
	fmt.Println(shortestCompletingWord("AN87005", []string{"participant", "individual", "start", "exist", "above", "already", "easy", "attack", "player", "important"}))
}

// leetcode748_最短完整词
func shortestCompletingWord(licensePlate string, words []string) string {
	m := make([]int, 26)
	licensePlate = strings.ToLower(licensePlate)
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			m[licensePlate[i]-'a']++
		}
	}
	var lists [16][]string
	for _, word := range words {
		lists[len(word)] = append(lists[len(word)], word)
	}
	for _, list := range lists {
		for _, word := range list {
			tempM := make([]int, 26)
			for i := 0; i < len(word); i++ {
				tempM[word[i]-'a']++
			}
			flag := true
			for k := range m {
				if tempM[k] < m[k] {
					flag = false
					break
				}
			}
			if flag == true {
				return word
			}
		}
	}
	return ""
}
