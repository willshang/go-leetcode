package main

import "fmt"

func main() {
	fmt.Println(minOperations(3))
}

func minOperations(n int) int {
	return n * n / 4
}
