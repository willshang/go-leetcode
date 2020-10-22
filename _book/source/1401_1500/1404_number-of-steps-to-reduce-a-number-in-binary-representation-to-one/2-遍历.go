package main

import "fmt"

func main() {
	fmt.Println(numSteps("1101"))
}

func numSteps(s string) int {
	res := 0
	flag := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			if flag == true {
				res = res + 2
			} else {
				res = res + 1 // 没有进位，遇0加1
			}
		} else {
			if flag == true {
				res++
			} else {
				if i != 0 {
					res = res + 2
				}
				flag = true
			}
		}
	}
	return res
}
