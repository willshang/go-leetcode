package main

import "math"

func main() {

}

var res int

func shoppingOffers(price []int, special [][]int, needs []int) int {
	res = math.MaxInt32
	dfs(price, special, needs, 0)
	return res
}

func dfs(price []int, special [][]int, needs []int, index int) {
	res := 0
	for i := 0; i < len(needs); i++ {
		res = res + needs[i]*price[i]
	}
}
