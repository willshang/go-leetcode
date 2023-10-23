package main

import (
	"fmt"
	"strings"
)

func main() {

}

// leetcode1016_子串能表示从1到N数字的二进制串
func queryString(s string, n int) bool {
	for i := 1; i <= n; i++ {
		target := fmt.Sprintf("%b", i)
		if strings.Contains(s, target) == false {
			return false
		}
	}
	return true
}
