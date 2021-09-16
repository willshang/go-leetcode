package main

import "sort"

func main() {

}

func maxmiumScore(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	res := 0
	sum := 0
	for i := 0; i < cnt; i++ {
		sum = sum + cards[i]
	}
	if sum%2 == 0 { // 偶数直接返回
		return sum
	}
	// 使用不同奇偶性的数替换：
	// 1、找个一个较小的数替换：cards[cnt-1]
	// 2、找到一个较小的数替换：跟cards[cnt-1]不同的较小的数
	for i := cnt; i < len(cards); i++ { // 情况1：在后面找1个数替换前cnt个最大数的最后1个数
		if cards[i]%2 != cards[cnt-1]%2 {
			res = max(res, sum-cards[cnt-1]+cards[i])
			break
		}
	}
	for i := cnt - 2; i >= 0; i-- { // 情况2：尝试找到1个奇偶性不同于cards[cnt-1]的数，然后替换掉
		if cards[i]%2 != cards[cnt-1]%2 {
			for j := cnt; j < len(cards); j++ {
				if cards[j]%2 != cards[i]%2 {
					res = max(res, sum-cards[i]+cards[j])
				}
			}
			break
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
