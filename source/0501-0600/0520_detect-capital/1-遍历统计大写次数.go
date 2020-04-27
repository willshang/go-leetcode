package main

import "fmt"

func main() {
	fmt.Println(detectCapitalUse("FlaG"))
	fmt.Println(detectCapitalUse("Flag"))
}

// leetcode520_检测大写字母
func detectCapitalUse(word string) bool {
	if word == "" {
		return false
	}
	count := 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			count++
		}
	}

	if count == 0 || count == len(word) ||
		(count == 1 && word[0] >= 'A' && word[0] <= 'Z') {
		return true
	}
	return false
}
