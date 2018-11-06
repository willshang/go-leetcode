package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "LL"
	fmt.Println(judgeCircle(str))
}
func judgeCircle(moves string) bool {
	return strings.Count(moves,"U") == strings.Count(moves,"D") &&
		strings.Count(moves,"L") == strings.Count(moves,"R")
}