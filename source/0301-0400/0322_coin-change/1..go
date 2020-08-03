package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)

}
