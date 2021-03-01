package main

import "fmt"

func main() {
	fmt.Println(minSteps("bab", "aba"))
}

func minSteps(s string, t string) int {
	res := 0
	a := make(map[int]int)
	b := make(map[int]int)
	for i := 0; i < len(s); i++ {
		a[int(s[i]-'a')]++
		b[int(t[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if a[i] < b[i] {
			res = res + b[i] - a[i]
		}
	}
	return res
}
