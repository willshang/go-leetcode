package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"}))
}

// leetcode1366_通过投票对团队排名
type Node struct {
	rank     []int
	teamName byte
}

func rankTeams(votes []string) string {
	m := make(map[byte][]int, 0)
	for i := 0; i < len(votes[0]); i++ {
		team := votes[0][i]
		m[team] = make([]int, len(votes[0]))
	}
	for _, vote := range votes {
		for i := 0; i < len(vote); i++ {
			m[vote[i]][i]++
		}
	}
	arr := make([]Node, 0)
	for k, v := range m {
		arr = append(arr, Node{
			rank:     v,
			teamName: k,
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		for k := 0; k < len(arr[i].rank); k++ {
			if arr[i].rank[k] != arr[j].rank[k] {
				return arr[i].rank[k] > arr[j].rank[k]
			}
		}
		return arr[i].teamName < arr[j].teamName
	})
	res := ""
	for i := 0; i < len(arr); i++ {
		res = res + string(arr[i].teamName)
	}
	return res
}
