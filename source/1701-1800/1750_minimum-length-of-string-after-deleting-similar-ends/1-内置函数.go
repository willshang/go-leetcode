package main

import "strings"

func main() {

}

func minimumLength(s string) int {
	for len(s) > 0 {
		if s[0] == s[len(s)-1] && len(s) != 1 {
			temp := string(s[0])
			s = strings.TrimLeft(s, temp)
			s = strings.TrimRight(s, temp)
		} else {
			break
		}
	}
	return len(s)
}
