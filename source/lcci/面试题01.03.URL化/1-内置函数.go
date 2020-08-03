package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpaces("Mr John Smith    ", 13))
}

func replaceSpaces(S string, length int) string {
	return strings.ReplaceAll(S[:length], " ", "%20")
}
