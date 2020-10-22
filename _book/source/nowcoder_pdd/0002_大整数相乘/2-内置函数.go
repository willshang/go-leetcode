package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b string
	for {
		n, _ := fmt.Scanf("%s %s", &a, &b)
		if n == 0 {
			break
		}
		fmt.Println(multiply(a, b))
	}
}

func multiply(num1 string, num2 string) string {
	a, b := new(big.Int), new(big.Int)
	a.SetString(num1, 10)
	b.SetString(num2, 10)
	a.Mul(a, b)
	return a.String()
}
