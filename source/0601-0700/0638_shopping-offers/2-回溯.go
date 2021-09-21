package main

import "fmt"

func main() {
	fmt.Println(shoppingOffers([]int{2, 5},
		[][]int{
			{3, 0, 5},
			{1, 2, 10},
		}, []int{3, 2}))
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	return dfs(price, special, needs)
}

func dfs(price []int, special [][]int, needs []int) int {
	res := 0
	for i := 0; i < len(needs); i++ { // 默认：走单品所需要的总价格
		res = res + needs[i]*price[i]
	}
	for i := 0; i < len(special); i++ { // 遍历每个礼包，每次取1份尝试
		j := 0
		for j = 0; j < len(needs); j++ {
			if needs[j] < special[i][j] { // 剪枝：不满足当前礼包要求，提前退出
				break
			}
		}
		if j == len(needs) { // 可以取该礼包，继续递归
			for k := 0; k < len(needs); k++ {
				needs[k] = needs[k] - special[i][k]
			}
			res = min(res, dfs(price, special, needs)+special[i][j]) // 递归，取最小
			for k := 0; k < len(needs); k++ {
				needs[k] = needs[k] + special[i][k]
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
