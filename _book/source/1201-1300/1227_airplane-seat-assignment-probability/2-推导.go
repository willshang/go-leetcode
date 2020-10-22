package main

import "fmt"

func main() {
	fmt.Println(nthPersonGetsNthSeat(10))
}

// leetcode1227_飞机座位分配概率
func nthPersonGetsNthSeat(n int) float64 {
	res := 1.0
	sum := 1.0
	// f(n) = 1/n * (f(n-1)+f(n-2)+f(n-3)+...+f(2)+1)
	for i := 2; i <= n; i++ {
		nth := 1.0 / float64(i)
		res = nth * sum
		sum += res
	}
	return res
}
