package main

func main() {

}

// leetcode2101_引爆最多的炸弹
var arr [][]int
var count int

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	arr = make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && judge(bombs[i][0], bombs[i][1], bombs[j][0], bombs[j][1], bombs[i][2]) == true {
				arr[i] = append(arr[i], j)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		count = 0
		dfs(i, make([]bool, n))
		if count > res {
			res = count
		}
	}
	return res
}

func dfs(start int, visited []bool) {
	visited[start] = true
	count++
	for i := 0; i < len(arr[start]); i++ {
		next := arr[start][i]
		if visited[next] == false {
			dfs(next, visited)
		}
	}
}

func judge(a, b, c, d, r int) bool {
	return r*r >= (a-c)*(a-c)+(b-d)*(b-d)
}
