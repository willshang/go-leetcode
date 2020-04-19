package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aaa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteArr := strings.Split(ransomNote, "")
	magazineArr := strings.Split(magazine, "")
	sort.Strings(ransomNoteArr)
	sort.Strings(magazineArr)

	i := 0
	j := 0
	for i < len(ransomNoteArr) && j < len(magazineArr) {
		if ransomNoteArr[i] > magazineArr[j] {
			j++
		} else if ransomNoteArr[i] < magazineArr[j] {
			return false
		} else {
			i++
			j++
		}
	}
	return i == len(ransomNote)
}
