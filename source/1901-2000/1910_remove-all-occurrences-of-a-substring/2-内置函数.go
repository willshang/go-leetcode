package main

import "strings"

func main() {

}

func removeOccurrences(s string, part string) string {
	for {
		index := strings.Index(s, part)
		if index < 0 {
			return s
		}
		s = s[:index] + s[index+len(part):]
	}
	return s
}
