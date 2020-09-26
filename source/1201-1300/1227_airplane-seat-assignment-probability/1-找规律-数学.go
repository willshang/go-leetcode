package main

import "fmt"

func main() {
	fmt.Println(nthPersonGetsNthSeat(10))
}

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}
