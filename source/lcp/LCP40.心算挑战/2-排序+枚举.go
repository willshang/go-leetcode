package main

import "sort"

func main() {

}

// leetcode_lcp40_心算挑战
func maxmiumScore(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	a, b := make([]int, 1), make([]int, 1)
	for i := 0; i < len(cards); i++ {
		if cards[i]%2 == 0 {
			a = append(a, cards[i]+a[len(a)-1])
		} else {
			b = append(b, cards[i]+b[len(b)-1])
		}
	}
	res := 0
	for i := 0; i <= cnt; i++ { // 枚举奇数的个数
		n, m := cnt-i, i
		if n < len(a) && m < len(b) && (a[n]+b[m])%2 == 0 {
			res = max(res, a[n]+b[m])
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
