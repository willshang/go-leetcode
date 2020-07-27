package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(15, 4))
}

func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}
