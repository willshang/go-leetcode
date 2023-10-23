package main

import "fmt"

func main() {
	fmt.Println(calculate("AB"))
}

func calculate(s string) int {
	return 1 << len(s)
}
