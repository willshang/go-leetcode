package main

import "sort"

func main() {

}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := make(map[int]int)
	for i := 0; i < len(items1); i++ {
		a, b := items1[i][0], items1[i][1]
		m[a] += b
	}
	for i := 0; i < len(items2); i++ {
		a, b := items2[i][0], items2[i][1]
		m[a] += b
	}
	res := make([][]int, 0)
	for k, v := range m {
		res = append(res, []int{k, v})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}
