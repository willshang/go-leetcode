package main

import (
	"fmt"
)

func main() {
	fmt.Println(addDigits(38))
	for i := 0; i < 1000; i++ {
		fmt.Printf("%4d => %4d\n", i, addDigits(i))
	}
}

func addDigits(num int) int {
	sum := 0
	for num != 0 {
		sum = sum + num%10
		num = num / 10
	}
	if sum/10 == 0 {
		return sum
	}
	return addDigits(sum)
}
