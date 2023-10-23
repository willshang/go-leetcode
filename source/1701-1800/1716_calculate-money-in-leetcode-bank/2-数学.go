package main

import "fmt"

func main() {
	// fmt.Println(totalMoney(4))
	fmt.Println(totalMoney(10))
}

func totalMoney(n int) int {
	res := 0
	a, b := n/7, n%7
	res = res + (28+(a-1)*7+28)*a/2
	res = res + (a+1+a+b)*b/2
	return res
}
