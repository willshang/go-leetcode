package main

import (
	"strings"
)

func main() {

}

func bestHand(ranks []int, suits []byte) string {
	if strings.Count(string(suits), string(suits[0])) == 5 {
		return "Flush"
	}
	m := make(map[int]int)
	for _, v := range ranks {
		m[v]++
		if m[v] == 3 {
			return "Three of a Kind"
		}
	}
	for _, v := range m {
		if v >= 2 {
			return "Pair"
		}
	}
	return "High Card"
}
