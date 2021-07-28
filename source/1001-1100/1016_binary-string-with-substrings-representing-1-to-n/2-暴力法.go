package main

import (
	"fmt"
	"strings"
)

func main() {

}

func queryString(s string, n int) bool {
	for i := n/2 + 1; i <= n; i++ {
		target := fmt.Sprintf("%b", i)
		if strings.Contains(s, target) == false {
			return false
		}
	}
	return true
}
