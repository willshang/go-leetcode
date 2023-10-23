package main

func main() {

}

// leetcode1672_最富有客户的资产总量
func maximumWealth(accounts [][]int) int {
	res := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum = sum + accounts[i][j]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
