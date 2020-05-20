package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(S string) string {
	res := ""
	begin := 1
	count := 1
	temp := ""
	for i := 0; i < len(S); i++ {
		if S[i] == ' ' {
			res = res + temp + strings.Repeat("a", count) + " "
			count++
			begin = 1
		} else {
			if begin == 1 {
				begin = 0
				if isBeginWithVowel(S[i]) {
					res = res + string(S[i])
					temp = "ma"
				} else {
					temp = string(S[i]) + "ma"
				}
			} else {
				res = res + string(S[i])
			}
		}
	}
	return res + temp + strings.Repeat("a", count)
}

func isBeginWithVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
