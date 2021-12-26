package main

func main() {

}

// leetcode2087_网格图中机器人回家的最小代价
func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	res := 0
	a, b, c, d := startPos[0], startPos[1], homePos[0], homePos[1]
	res = res - rowCosts[a] - colCosts[b]
	if a > c {
		a, c = c, a
	}
	if b > d {
		b, d = d, b
	}
	for i := a; i <= c; i++ {
		res = res + rowCosts[i]
	}
	for i := b; i <= d; i++ {
		res = res + colCosts[i]
	}
	return res
}
