package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkOnesSegment("1100"))
}

func checkOnesSegment(s string) bool {
	arr := strings.Split(s, "0")
	for i := 1; i < len(arr); i++ {
		if len(arr[i]) > 0 {
			return false
		}
	}
	return true
}
