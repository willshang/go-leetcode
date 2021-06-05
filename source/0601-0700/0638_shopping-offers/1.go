package main

func main() {

}

var res int

func shoppingOffers(price []int, special [][]int, needs []int) int {
	res = 0
	for i := 0; i < len(needs); i++ {
		res = res + needs[i]*price[i] // 全部单品需要的总价格
	}
	dfs(price, special, needs, 0)
	return res
}

func dfs(price []int, special [][]int, needs []int, index int) {

}
