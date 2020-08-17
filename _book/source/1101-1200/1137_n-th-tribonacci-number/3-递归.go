package main

import "fmt"

func main() {
	fmt.Println(tribonacci(5))
}

var m map[int]int

func tribonacci(n int) int {
	if m == nil {
		m = make(map[int]int)
	}
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if value, ok := m[n]; ok {
		return value
	} else {
		value := tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
		m[n] = value
	}
	return m[n]
}
