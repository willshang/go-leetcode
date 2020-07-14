package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	for {
		n, _ := fmt.Scanf("%f", &num)
		if n == 0 {
			break
		}
		fmt.Println(int(math.Floor(num + 0.5)))
	}
}
