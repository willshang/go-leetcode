package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
	m := make(map[string]string)
	for i := 10; i <= 26; i++ {
		m[strconv.Itoa(i)+"#"] = string('j' + i - 10)
	}
	m2 := make(map[string]string)
	for i := 1; i <= 9; i++ {
		m2[strconv.Itoa(i)] = string('a' + i - 1)
	}
	for k, v := range m {
		s = strings.ReplaceAll(s, k, v)
	}
	for k, v := range m2 {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}
