package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	count := 0
	if s[0] == 'L' {
		count++
	} else {
		count--
	}
	for i := 1; i < len(s); i++ {
		if s[i] == 'L' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}
