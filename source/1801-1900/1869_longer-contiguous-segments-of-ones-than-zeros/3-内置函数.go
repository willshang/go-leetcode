package main

import "strings"

func main() {

}

func checkZeroOnes(s string) bool {
	a, b := 0, 0
	for _, v := range strings.Split(s, "1") {
		b = max(b, len(v))
	}
	for _, v := range strings.Split(s, "0") {
		a = max(a, len(v))
	}
	return b > a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
