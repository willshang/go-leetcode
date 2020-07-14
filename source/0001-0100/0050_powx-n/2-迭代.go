package main

import "fmt"

func main() {
	fmt.Println(myPow(2.00000, 10))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := float64(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * x
		}
		x = x * x
		n = n / 2
	}
	return res
}
