package main

import "fmt"

func main() {
	fmt.Println(numWays(100))
}

var m = make(map[int]int)

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if m[n] > 0 {
		return m[n]
	} else {
		m[n] = (numWays(n-1) + numWays(n-2)) % 1000000007
	}
	return m[n]
}
