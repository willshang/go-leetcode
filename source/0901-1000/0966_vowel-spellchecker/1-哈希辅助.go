package main

import (
	"fmt"
	"strings"
)

func main() {
	//a := []string{"KiTe", "kite", "hare", "Hare"}
	//b := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}

	a := []string{"ae", "aa"}
	b := []string{"UU"}
	fmt.Println(spellchecker(a, b))
}

// leetcode966_元音拼写检查器
func spellchecker(wordlist []string, queries []string) []string {
	m := make(map[string]bool)
	mLower := make(map[string]string)
	mVowel := make(map[string]string)
	for i := 0; i < len(wordlist); i++ {
		m[wordlist[i]] = true
		lowerStr := strings.ToLower(wordlist[i])
		if _, ok := mLower[lowerStr]; ok == false {
			mLower[lowerStr] = wordlist[i]
		}
		vowelStr := changeVowel(lowerStr)
		if _, ok := mVowel[vowelStr]; ok == false {
			mVowel[vowelStr] = wordlist[i]
		}
	}
	for i := 0; i < len(queries); i++ {
		if m[queries[i]] == true {
			continue
		}
		lowerStr := strings.ToLower(queries[i])
		if v, ok := mLower[lowerStr]; ok {
			queries[i] = v
			continue
		}
		vowelStr := changeVowel(lowerStr)
		if v, ok := mVowel[vowelStr]; ok {
			queries[i] = v
			continue
		}
		queries[i] = ""
	}
	return queries
}

func changeVowel(str string) string {
	str = strings.ReplaceAll(str, "a", "?")
	str = strings.ReplaceAll(str, "e", "?")
	str = strings.ReplaceAll(str, "i", "?")
	str = strings.ReplaceAll(str, "o", "?")
	str = strings.ReplaceAll(str, "u", "?")
	return str
}
