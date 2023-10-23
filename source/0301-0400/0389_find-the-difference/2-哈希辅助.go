package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	bytest := []byte(t)
	bytess := []byte(s)
	for _, v := range bytest {
		m[v]++
	}
	for _, v := range bytess {
		m[v]--
	}
	for k := range m {
		if m[k] == 1 {
			return k
		}
	}
	return 0
}
