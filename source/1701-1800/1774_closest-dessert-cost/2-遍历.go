package main

import "math"

func main() {

}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	res := math.MaxInt32
	n, m := len(baseCosts), len(toppingCosts)
	for i := 0; i < n; i++ {
		for j := 0; j < (1 << m); j++ { // 选择第1次
			for k := j; k < (1 << m); k++ { // 选择第2次
				total := baseCosts[i]
				for l := 0; l < m; l++ {
					if j&(1<<l) != 0 {
						total = total + toppingCosts[l]
					}
					if k&(1<<l) != 0 {
						total = total + toppingCosts[l]
					}
				}
				if abs(res-target) > abs(total-target) {
					res = total
				} else if abs(res-target) == abs(total-target) && total < res {
					res = total
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
