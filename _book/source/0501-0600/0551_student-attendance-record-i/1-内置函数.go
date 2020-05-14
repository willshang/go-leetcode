package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}

func checkRecord(s string) bool {
	if strings.Count(s, "A") <= 1 && strings.Count(s, "LLL") == 0 {
		return true
	}
	return false
}
