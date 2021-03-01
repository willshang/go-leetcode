package main

import "sort"

func main() {

}

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	sort.Ints(deck)
	res := make([]int, n)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = i
	}
	for i := 0; i < n; i++ {
		first := temp[0]
		temp = temp[1:]
		res[first] = deck[i]
		if len(temp) > 0 {
			first := temp[0]
			temp = temp[1:]
			temp = append(temp, first)
		}
	}
	return res
}
