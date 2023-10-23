package main

import "fmt"

func main() {
	fmt.Println(countVowelSubstrings("aeiouu"))
}

// leetcode2062_统计字符串中的元音子字符串
var m = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func countVowelSubstrings(word string) int {
	n := len(word)
	res := 0
	for i := 0; i < n; i++ {
		temp := make(map[byte]bool)
		for j := i; j < n; j++ {
			if m[word[j]] == false {
				break
			}
			temp[word[j]] = true
			if len(temp) == 5 {
				res++
			}
		}
	}
	return res
}
