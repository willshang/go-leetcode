package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(countValidWords("cat and dog"))
}

// leetcode2047_句子中的有效单词数
var m map[byte]bool = map[byte]bool{
	'!': true,
	'.': true,
	',': true,
}

func countValidWords(sentence string) int {
	res := 0
	arr := strings.Fields(sentence)
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if m[temp[len(temp)-1]] == true {
			temp = temp[:len(temp)-1]
		}
		if strings.ContainsAny(arr[i], "0123456789") ||
			strings.ContainsAny(temp, "!.,") ||
			strings.Count(temp, "-") >= 2 {
			continue
		}
		j := strings.IndexByte(temp, '-')
		// 存在-
		if j >= 0 && (j == 0 || j == len(temp)-1 || unicode.IsLower(rune(temp[j-1])) == false ||
			unicode.IsLower(rune(temp[j+1])) == false) {
			continue
		}
		res++
	}
	return res
}
