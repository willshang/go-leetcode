package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 2},
		{1, 2},
		{3, 3},
		{1, 5},
		{1, 5},
	}
	fmt.Println(maxEvents(arr))
}

func maxEvents(events [][]int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := events[i][0]; j <= events[i][1]; j++ {
			if m[j] == false {
				m[j] = true
				break
			}
		}
	}
	fmt.Println(m)
	return len(m)
}
