package main

import "fmt"

func main() {
	fmt.Println(add(3, 5))
}

func add(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, (a&b)<<1)
}
