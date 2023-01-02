package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := getResult(n)
	fmt.Println(res)
}

// res = n^n-n+1
func getResult(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * n
	}
	return res - n + 1
}
