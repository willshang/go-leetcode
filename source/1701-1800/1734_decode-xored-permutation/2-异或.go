package main

import "fmt"

func main() {
	fmt.Println(decode([]int{3, 1}))
}

func decode(encoded []int) []int {
	n := len(encoded)
	first := 0
	for i := 1; i <= n+1; i++ {
		first = first ^ i
	}
	for i := 1; i <= n; i = i + 2 {
		first = first ^ encoded[i]
	}
	res := make([]int, n+1)
	res[0] = first
	for i := 0; i < n; i++ {
		res[i+1] = encoded[i] ^ res[i]
	}
	return res
}
