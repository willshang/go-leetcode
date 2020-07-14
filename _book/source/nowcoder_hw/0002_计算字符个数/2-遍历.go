package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var c uint8
	fmt.Scanf("%s", &str)
	fmt.Scanf("%c", &c)
	str = strings.ToLower(str)
	c = strings.ToLower(string(c))[0]
	res := 0
	for i := 0; i < len(str); i++ {
		if str[i] == uint8(c) {
			res++
		}
	}
	fmt.Println(res)
	return
}
