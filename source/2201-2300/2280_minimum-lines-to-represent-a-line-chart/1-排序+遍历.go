package main

import "sort"

func main() {

}

// leetcode2280_表示一个折线图的最少线段数
func minimumLines(stockPrices [][]int) int {
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})
	res := 0
	dx, dy := 0, 0
	for i := 1; i < len(stockPrices); i++ {
		x := stockPrices[i][0] - stockPrices[i-1][0]
		y := stockPrices[i][1] - stockPrices[i-1][1]
		if i == 1 {
			dx, dy = x, y
			res++
		} else if dx*y != x*dy { // y/x = dy/dx =>  dx*y == x*dy
			dx, dy = x, y
			res++
		}
	}
	return res
}
