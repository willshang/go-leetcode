package main

func main() {

}

// leetcode851_喧闹和富有
var res []int

func loudAndRich(richer [][]int, quiet []int) []int {
	arr := make(map[int][]int)
	n := len(quiet)
	res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	for i := 0; i < len(richer); i++ { // 有向无环图
		a, b := richer[i][0], richer[i][1]
		arr[b] = append(arr[b], a) // a比b更有钱
	}
	for i := 0; i < n; i++ {
		dfs(quiet, arr, i)
	}
	return res
}

func dfs(quiet []int, arr map[int][]int, i int) int {
	if res[i] == -1 {
		res[i] = i // 首先最安静的人等于本身
		for j := 0; j < len(arr[i]); j++ {
			next := dfs(quiet, arr, arr[i][j]) // 递归找到arr[i][j]对应的最安静的人
			if quiet[next] < quiet[res[i]] {
				res[i] = next
			}
		}
	}
	return res[i]
}
