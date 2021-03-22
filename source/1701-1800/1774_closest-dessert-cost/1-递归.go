package main

import "math"

func main() {

}

// leetcode1774_最接近目标价格的甜点成本
var res int

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	res = math.MaxInt32
	for i := 0; i < len(baseCosts); i++ {
		dfs(toppingCosts, target, baseCosts[i], 0)
	}
	return res
}

func dfs(toppingCosts []int, target int, sum int, index int) {
	if abs(res-target) > abs(sum-target) {
		res = sum
	} else if abs(res-target) == abs(sum-target) && sum < res {
		res = sum
	}
	if sum > target {
		return
	}
	if index == len(toppingCosts) {
		return
	}
	dfs(toppingCosts, target, sum, index+1)
	dfs(toppingCosts, target, sum+toppingCosts[index], index+1)
	dfs(toppingCosts, target, sum+2*toppingCosts[index], index+1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
