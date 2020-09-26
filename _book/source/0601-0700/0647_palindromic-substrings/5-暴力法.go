package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaaaa"))
}

func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	res := len(s)
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && judge(s, i, j) == true {
				res++
			}
		}
	}
	return res
}

func judge(s string, i, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
