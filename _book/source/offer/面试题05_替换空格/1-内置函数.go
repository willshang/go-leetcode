package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
	// return strings.ReplaceAll(s, " ", "%20")
}
