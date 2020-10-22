package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}

func maxDepth(s string) int {
	res := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			if count > res {
				res = count
			}
			count--
		}
	}
	return res
}
