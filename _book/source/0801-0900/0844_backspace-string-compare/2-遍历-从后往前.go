package main

import (
	"fmt"
)

func main() {
	S := "a##c"
	T := "#a#c"
	fmt.Println(backspaceCompare(S, T))
}

func backspaceCompare(S string, T string) bool {
	return check(S) == check(T)
}

func check(S string) string {
	str := ""
	count := 0
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '#' {
			count++
		} else {
			if count != 0 {
				count--
				continue
			}
			str = string(S[i]) + str
		}
	}
	return str
}
