package main

import (
	"sort"
)

func main() {

}

func findWinners(matches [][]int) [][]int {
	m := make(map[int]int)
	for i := 0; i < len(matches); i++ {
		a, b := matches[i][0], matches[i][1]
		if _, ok := m[a]; ok == false {
			m[a] = 0 // 赢的标记为0
		}
		m[b]++ // 输的+1
	}
	arr := [2][]int{}
	for k, v := range m {
		if v < 2 { // 0：全赢， 1：输1次
			arr[v] = append(arr[v], k)
		}
	}
	sort.Ints(arr[0])
	sort.Ints(arr[1])
	return [][]int{arr[0], arr[1]}
}
