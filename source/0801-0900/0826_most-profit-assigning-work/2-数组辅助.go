package main

func main() {

}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	res := 0
	arr := make([]int, 100001) // 难度对应的最大利润
	for i := 0; i < len(difficulty); i++ {
		arr[difficulty[i]] = max(arr[difficulty[i]], profit[i])
	}
	maxProfit := arr[0]
	for i := 1; i < 100001; i++ {
		maxProfit = max(maxProfit, arr[i])
		arr[i] = maxProfit
	}
	for i := 0; i < len(worker); i++ {
		res = res + arr[worker[i]]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
