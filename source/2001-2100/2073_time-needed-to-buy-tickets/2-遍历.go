package main

func main() {

}

// leetcode2073_买票需要的时间
func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			res = res + min(tickets[k], tickets[i]) // 前面的人，最多跟第k个人一样多
		} else {
			res = res + min(tickets[k]-1, tickets[i]) // 后面的人，少一次选择
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
