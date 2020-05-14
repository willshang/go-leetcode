package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(str))
}

var qRow = map[byte]bool{
	'q': true,
	'w': true,
	'e': true,
	'r': true,
	't': true,
	'y': true,
	'u': true,
	'i': true,
	'o': true,
	'p': true,
}

var aRow = map[byte]bool{
	'a': true,
	's': true,
	'd': true,
	'f': true,
	'g': true,
	'h': true,
	'j': true,
	'k': true,
	'l': true,
}

var zRow = map[byte]bool{
	'z': true,
	'x': true,
	'c': true,
	'v': true,
	'b': true,
	'n': true,
	'm': true,
}

func findWords(words []string) []string {
	res := make([]string, 0, len(words))
	for _, word := range words {
		w := strings.ToLower(word)
		if isAllIn(w, qRow) || isAllIn(w, aRow) || isAllIn(w, zRow) {
			res = append(res, word)
		}
	}
	return res
}

func isAllIn(s string, Row map[byte]bool) bool {
	for i := range s {
		if !Row[s[i]] {
			return false
		}
	}
	return true
}
