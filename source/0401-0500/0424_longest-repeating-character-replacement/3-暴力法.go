package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
}

func characterReplacement(s string, k int) int {
	if s == "" {
		return 0
	}
	res := 0
	for i := 0; i < len(s); i++ {
		temp := k
		j := i + 1
		for ; j < len(s); j++ {
			if s[j] != s[i] {
				if temp == 0 {
					break
				}
				temp--
			}
		}
		if j-i+temp > len(s) {
			return len(s)
		}
		if j-i+temp > res {
			res = j - i + temp
		}
	}
	return res
}
