package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(detectCapitalUse("FlaG"))
	fmt.Println(detectCapitalUse("Flag"))
}

func detectCapitalUse(word string) bool {
	pattern := "(^[a-z]+)$|(^[A-Z]+)$|(^[A-Z]{1}[a-z]*)$"
	isMatch, _ := regexp.MatchString(pattern, word)
	return isMatch
}
