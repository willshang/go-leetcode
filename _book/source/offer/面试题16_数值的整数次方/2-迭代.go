package main

import "fmt"

func main() {
	fmt.Println(myPow(float64(2.0), -2))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	res := 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n >= 1 {
		if n%2 == 1 {
			res = res * x
			n--
		} else {
			x = x * x
			n = n / 2
		}
	}
	return res
}
