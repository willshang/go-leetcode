package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minMutation("AAAAAAAA", "CCCCCCCC", []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC",
		"AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA", "CCCCCCCC"}))
}

var res int

func minMutation(start string, end string, bank []string) int {
	res = math.MaxInt32
	m := make(map[string]bool)
	for i := 0; i < len(bank); i++ {
		m[bank[i]] = true
	}
	if _, ok := m[end]; ok == false {
		return -1
	}
	dfs(start, end, 0, bank, make([]bool, len(bank)))
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func dfs(start string, end string, index int, bank []string, visited []bool) {
	if start == end {
		if index < res {
			res = index
		}
		return
	}
	for i := 0; i < len(bank); i++ {
		if visited[i] == false && judge(start, bank[i]) {
			visited[i] = true
			dfs(bank[i], end, index+1, bank, visited)
			visited[i] = false
		}
	}
}

func judge(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count == 1
}
