package main

func main() {

}

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	res := 0
	a, b, c, d := startPos[0], startPos[1], homePos[0], homePos[1]
	if a > c {
		for i := a - 1; i >= c; i-- {
			res = res + rowCosts[i]
		}
	} else if a < c {
		for i := a + 1; i <= c; i++ {
			res = res + rowCosts[i]
		}
	}
	if b > d {
		for i := b - 1; i >= d; i-- {
			res = res + colCosts[i]
		}
	} else if b < d {
		for i := b + 1; i <= d; i++ {
			res = res + colCosts[i]
		}
	}
	return res
}
