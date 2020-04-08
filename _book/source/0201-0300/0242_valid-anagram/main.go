package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	rec := make(map[rune]int, len(sr))
	for i := range sr {
		rec[sr[i]]++
		rec[tr[i]]--
	}

	for _, n := range rec {
		if n != 0 {
			return false
		}
	}
	return true
}
