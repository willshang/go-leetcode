package main

func main() {

}

// leetcode2147_分隔长廊的方案数
var mod = 1000000007

func numberOfWays(corridor string) int {
	res := 1
	prev := -1
	count := 0
	for i := 0; i < len(corridor); i++ {
		if corridor[i] == 'S' {
			count++
			// 统计每2个1组之间有多少可能；相乘即可
			if count >= 3 && count%2 == 1 { // S>=3+奇数
				res = res * (i - prev) % mod
			}
			prev = i // 上一个座位的位置
		}
	}
	if count < 2 || count%2 == 1 {
		return 0
	}
	return res
}
