package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
}

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	return len(strings.Fields(s))
}
