package main

import "fmt"

func main() {
	fmt.Println(divisorGame(2))
}

func divisorGame(N int) bool {
	return N%2 == 0
}
