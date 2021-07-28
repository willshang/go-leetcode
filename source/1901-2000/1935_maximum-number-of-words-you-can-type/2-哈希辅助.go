package main

import "strings"

func main() {

}

func canBeTypedWords(text string, brokenLetters string) int {
	res := 0
	m := make(map[byte]bool)
	for i := 0; i < len(brokenLetters); i++ {
		m[brokenLetters[i]] = true
	}
	arr := strings.Split(text, " ")
	for i := 0; i < len(arr); i++ {
		flag := true
		for j := 0; j < len(arr[i]); j++ {
			if m[arr[i][j]] == true {
				flag = false
				break
			}
		}
		if flag == true {
			res++
		}
	}
	return res
}
