package main

import "strings"

func main() {

}

func makeFancyString(s string) string {
	for i := 'a'; i <= 'z'; i++ {
		for strings.Contains(s, strings.Repeat(string(i), 3)) == true {
			s = strings.ReplaceAll(s, strings.Repeat(string(i), 3), strings.Repeat(string(i), 2))
		}
	}
	return s
}
