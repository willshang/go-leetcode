package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}
func toGoatLatin(S string) string {
	ss := strings.Split(S, " ")
	for i := range ss {
		ss[i] = handle(ss[i], i)
	}
	return strings.Join(ss, " ")
}

func handle(s string, i int) string {
	postfix := "ma" + strings.Repeat("a", i+1)
	if isBeginWithVowel(s) {
		return s + postfix
	}
	return s[1:] + s[0:1] + postfix
}

func isBeginWithVowel(s string) bool {
	switch s[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
