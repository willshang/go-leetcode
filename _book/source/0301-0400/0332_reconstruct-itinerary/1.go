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

var m map[string][]string
var res []string

func findItinerary(tickets [][]string) []string {
	res = make([]string, 0)
	m = make(map[string][]string)
	for i := 0; i < len(tickets); i++ {
		m[tickets[i][0]] = append(m[tickets[i][0]], tickets[i][1])
	}
	for _, v := range m {
		sort.Strings(v)
	}
	dfs("JFK")
	return res
}

func dfs(start string) {
	for _, v := range m[start] {

	}
}
