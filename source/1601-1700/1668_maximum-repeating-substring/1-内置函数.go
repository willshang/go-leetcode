package main

import "strings"

func main() {

}

// leetcode1668_最大重复子字符串
func maxRepeating(sequence string, word string) int {
	res := 0
	for i := 1; ; i++ {
		if strings.Contains(sequence, strings.Repeat(word, i)) == false {
			break
		} else {
			res = i
		}
	}
	return res
}
