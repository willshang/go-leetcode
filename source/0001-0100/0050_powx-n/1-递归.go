package main

import "fmt"

func main() {
	fmt.Println(myPow(2.00000, 10))
}

// leetcode50_Pow(x,n)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 1 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}
