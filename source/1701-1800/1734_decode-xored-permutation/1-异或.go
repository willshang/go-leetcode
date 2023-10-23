package main

import "fmt"

func main() {
	fmt.Println(decode([]int{3, 1}))
}

// leetcode1734_解码异或后的排列
func decode(encoded []int) []int {
	n := len(encoded)
	temp := make([]int, n)
	copy(temp, encoded)
	first := encoded[0]
	for i := 1; i < n; i++ {
		temp[i] = encoded[i] ^ temp[i-1]
		first = first ^ temp[i]
	}
	for i := 1; i <= n+1; i++ {
		first = first ^ i
	}

	res := make([]int, n+1)
	res[0] = first
	for i := 0; i < n; i++ {
		res[i+1] = encoded[i] ^ res[i]
	}
	return res
}
