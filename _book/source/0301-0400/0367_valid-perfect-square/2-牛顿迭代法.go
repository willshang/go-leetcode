package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(100))
	fmt.Println(isPerfectSquare(99))
}

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	x := num / 2
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}
