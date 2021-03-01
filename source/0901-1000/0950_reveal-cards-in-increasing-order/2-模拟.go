package main

import "sort"

func main() {

}

// leetcode950_按递增顺序显示卡牌
func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	sort.Ints(deck)
	if n <= 2 {
		return deck
	}
	res := make([]int, 0)
	res = append(res, deck[n-1])
	for i := n - 2; i >= 0; i-- {
		// 插入一个数a，要将序列尾部的数b拿出来，组成[a,b]放入序列头部，形成新序列。
		last := res[len(res)-1]
		res = res[:len(res)-1]
		res = append([]int{deck[i], last}, res...)
	}
	return res
}
