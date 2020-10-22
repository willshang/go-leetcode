package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(15, 4))
}

// leetcode1518_换酒问题
func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles > 0 {
		times := numBottles / numExchange
		res = res + times
		numBottles = numBottles%numExchange + times
		if numBottles < numExchange {
			break
		}
	}
	return res
}
