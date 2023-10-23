package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(splitString("54321"))
}

func splitString(s string) bool {
	n := len(s)
	for i := 0; i < n-1; i++ {
		a, _ := strconv.Atoi(s[0 : i+1])
		index := i
		for j := i + 1; j < n; j++ {
			b, _ := strconv.Atoi(s[index+1 : j+1])
			c, _ := strconv.Atoi(s[index+1 : n])
			if c == a-1 {
				return true
			}
			if b == a-1 {
				index = j
				a = b
				c, _ := strconv.Atoi(s[index+1 : n])
				if c == a-1 {
					return true
				}
			} else if b > a-1 {
				break
			}
		}
	}
	return false
}
