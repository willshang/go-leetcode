package main

import "sort"

func main() {

}

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	sort.Ints(deck)
	if n <= 2 {
		return deck
	}
	res := make([]int, 0)
	res = append(res, deck[n-1])
	for i := n - 2; i >= 0; i-- {
		res = append(res, deck[i])
		first := res[0]
		res = res[1:]
		res = append(res, first)
	}
	last := res[len(res)-1]
	res = res[:len(res)-1]
	res = append([]int{last}, res...)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return res
}
