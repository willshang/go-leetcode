package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateTheString(2))
}

func generateTheString(n int) string {
	res := ""
	if n%2 == 0 {
		res = "a"
		for i := 0; i < n-1; i++ {
			res = res + "b"
		}
	} else {
		for i := 0; i < n; i++ {
			res = res + "a"
		}
	}
	return res
}
