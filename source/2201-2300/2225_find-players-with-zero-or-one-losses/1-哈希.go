package main

import "sort"

func main() {

}

// leetcode2225_找出输掉零场或一场比赛的玩家
func findWinners(matches [][]int) [][]int {
	m1, m2 := make(map[int]int), make(map[int]int)
	for i := 0; i < len(matches); i++ {
		a, b := matches[i][0], matches[i][1]
		m1[a]++ // 赢的+1
		m2[b]-- // 输的-1
	}
	arr1, arr2 := make([]int, 0), make([]int, 0)
	for k := range m1 {
		if _, ok := m2[k]; ok == false {
			arr1 = append(arr1, k)
		}
	}
	for k, v := range m2 {
		if v == -1 {
			arr2 = append(arr2, k)
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	return [][]int{arr1, arr2}
}
