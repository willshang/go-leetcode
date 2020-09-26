package main

import "fmt"

func main() {
	fmt.Println(swapNumbers([]int{1, 2}))
}

func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] ^ numbers[1]
	numbers[1] = numbers[1] ^ numbers[0]
	numbers[0] = numbers[0] ^ numbers[1]
	return numbers
}
