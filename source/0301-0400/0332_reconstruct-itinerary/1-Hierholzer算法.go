package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
		{"A", "C"},
		{"A", "D"},
		{"A", "E"},
		{"A", "B"},
	}
	fmt.Println(findItinerary(arr))
}

// leetcode332_重新安排行程
var m map[string][]string
var res []string

func findItinerary(tickets [][]string) []string {
	res = make([]string, 0)
	m = make(map[string][]string)
	for i := 0; i < len(tickets); i++ {
		a, b := tickets[i][0], tickets[i][1]
		m[a] = append(m[a], b)
	}
	for _, v := range m {
		sort.Strings(v)
	}
	dfs("JFK")
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

func dfs(start string) {
	for len(m[start]) > 0 {
		node := m[start][0]
		m[start] = m[start][1:]
		dfs(node)
	}
	res = append(res, start)
}
