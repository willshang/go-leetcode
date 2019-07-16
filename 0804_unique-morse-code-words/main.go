package main

import (
	"fmt"
)

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

var table = []string{".-", "-...", "-.-.", "-..", ".", "..-.",
	"--.", "....", "..", ".---", "-.-", ".-..", "--", "-.",
	"---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-",
	".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	res := map[string]bool{}
	for _, w := range words {
		b := ""
		for i := 0; i < len(w); i++ {
			b = b + table[w[i]-'a']
		}
		res[b] = true
	}
	return len(res)
}
