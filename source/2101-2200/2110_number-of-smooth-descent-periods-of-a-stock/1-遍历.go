package main

func main() {

}

// leetcode2110_股票平滑下跌阶段的数目
func getDescentPeriods(prices []int) int64 {
	res := int64(1)
	if len(prices) == 1 {
		return 1
	}
	count := int64(1)
	for i := 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] == 1 {
			count++
		} else {
			count = int64(1)
		}
		res = res + count
	}
	return res
}
