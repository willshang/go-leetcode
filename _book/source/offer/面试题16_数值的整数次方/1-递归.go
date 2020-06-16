package main

import "fmt"

func main() {
	fmt.Println(myPow(float64(2.0), 10))
}

// 剑指offer_面试题16_数值的整数次方
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	res := 1.0
	if n > 0 {
		res = myPow(x, n/2)
		return res * res * myPow(x, n%2)
	} else {
		res = myPow(x, -n/2)
		res = res * res * myPow(x, -n%2)
		return 1 / res
	}
}
