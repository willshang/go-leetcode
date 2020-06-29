package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var c string
	fmt.Scan(&str)
	fmt.Scan(&c)
	str = strings.ToLower(str)
	c = strings.ToLower(c)
	res := 0
	for i := 0; i < len(str); i++ {
		if strings.EqualFold(string(str[i]), c) {
			res++
		}
	}
	fmt.Println(res)
	return
}
