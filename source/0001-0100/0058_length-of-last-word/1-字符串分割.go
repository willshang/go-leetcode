package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a "))
}

func lengthOfLastWord(s string) int {
	arr := strings.Split(strings.Trim(s, " "), " ")
	return len(arr[len(arr)-1])
}
