package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; {
		if s[i] == '#' {
			value, _ := strconv.Atoi(string(s[i-2 : i]))
			res = string('a'+value-1) + res
			i = i - 3
		} else {
			value, _ := strconv.Atoi(string(s[i]))
			res = string('a'+value-1) + res
			i = i - 1
		}
	}
	return res
}
