package main

import (
	"fmt"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

func defangIPaddr(address string) string {
	res := ""
	for i := range address {
		if address[i] == '.' {
			res = res + "[.]"
		} else {
			res = res + string(address[i])
		}
	}
	return res
}
