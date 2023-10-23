package main

import "strings"

func main() {

}

// leetcode1935_可以输入的最大单词数
func canBeTypedWords(text string, brokenLetters string) int {
	res := 0
	arr := strings.Split(text, " ")
	for i := 0; i < len(arr); i++ {
		if strings.ContainsAny(arr[i], brokenLetters) == false {
			res++
		}
	}
	return res
}
