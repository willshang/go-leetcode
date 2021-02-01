package main

func main() {

}

// leetcode1443_收集树上所有苹果的最少时间
func minTime(n int, edges [][]int, hasApple []bool) int {
	arr := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	visited := make([]bool, n)
	res, _ := dfs(arr, hasApple, visited, 0)
	if res >= 2 {
		return res - 2 // 遍历N个点，长度：2N-2
	}
	return 0
}

func dfs(arr [][]int, hasApple, visited []bool, index int) (int, bool) {
	visited[index] = true
	res := 0
	has := false
	if hasApple[index] == true {
		has = true
	}
	for i := 0; i < len(arr[index]); i++ {
		next := arr[index][i]
		if visited[next] == true {
			continue
		}
		total, isExist := dfs(arr, hasApple, visited, next)
		if isExist {
			has = true
			res = res + total
		}
	}
	if has == true {
		return res + 2, true
	}
	return 0, false
}
