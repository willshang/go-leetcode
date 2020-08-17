package main

import "fmt"

func main() {
	fmt.Println(isUnique("leetcode"))
}

func isUnique(astr string) bool {
	for i := 0; i < len(astr); i++ {
		for j := i + 1; j < len(astr); j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}
