package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	c := 1
	for i := 1; i <= n; i++ {
		c = c * (n + i) / i
	}
	return c / (n + 1)
}
