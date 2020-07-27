package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

// leetcode389_找不同
func findTheDifference(s string, t string) byte {
	m := [26]int{}
	bytest := []byte(t)
	bytess := []byte(s)
	for _, v := range bytest {
		m[v-'a']++
	}
	for _, v := range bytess {
		m[v-'a']--
	}
	for k, _ := range m {
		if m[k] == 1 {
			return byte(k + 'a')
		}
	}
	return 0
}
