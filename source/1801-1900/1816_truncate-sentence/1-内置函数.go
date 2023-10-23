package main

import "strings"

func main() {

}

func truncateSentence(s string, k int) string {
	arr := strings.Split(s, " ")
	if k < len(arr) {
		return strings.Join(arr[:k], " ")
	}
	return s
}
