package main

import "fmt"

func main() {
	//fmt.Println(countVowelSubstrings("aeiouu"))
	fmt.Println(countVowelSubstrings("cuaieuouac"))
}

var m = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func countVowelSubstrings(word string) int {
	n := len(word)
	res := 0
	temp := make(map[byte]int)
	count := 1
	start := -1
	for i := 0; i < n; i++ {
		if m[word[i]] == false {
			temp = make(map[byte]int)
			count = 1
			start = i
			continue
		}
		temp[word[i]]++
		for temp[word[start+count]] > 1 { // 左边多余的字母个数
			temp[word[start+count]]--
			count++
		}
		if len(temp) == 5 {
			res = res + count
		}
	}
	return res
}
