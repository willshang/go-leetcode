package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(100))
	fmt.Println(isPerfectSquare(99))
}

func isPerfectSquare(num int) bool {
	i := 1
	for i*i < num {
		i++
	}
	return i*i == num
}
