package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(100))
	fmt.Println(isPerfectSquare(99))
}
func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num = num - i
		i = i + 2
	}
	return num == 0
}
