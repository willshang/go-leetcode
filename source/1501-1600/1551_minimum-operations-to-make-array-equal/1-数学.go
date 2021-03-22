package main

import "fmt"

func main() {
	fmt.Println(minOperations(3))
}

func minOperations(n int) int {
	res := 0
	for i := 1; i < n; i = i + 2 {
		res = res + n - i
	}
	return res
}
