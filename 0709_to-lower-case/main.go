package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	fmt.Println(toLowerCase(str))
}
func toLowerCase(str string) string {
	return strings.ToLower(str)
}