package main

import (
	"fmt"
)

func main() {
	fmt.Println(translateNum(10022))
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	var res int
	if num%100 <= 25 && num%100 > 9 {
		res = res + translateNum(num/100)
		res = res + translateNum(num/10)
	} else {
		res = res + translateNum(num/10)
	}
	return res
}
