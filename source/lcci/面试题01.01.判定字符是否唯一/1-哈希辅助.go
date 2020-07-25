package main

import "fmt"

func main() {
	fmt.Println(isUnique("leetcode"))
}

func isUnique(astr string) bool {
	m := make(map[byte]bool)
	for i := 0; i < len(astr); i++ {
		if m[astr[i]] == true {
			return false
		}
		m[astr[i]] = true
	}
	return true
}
