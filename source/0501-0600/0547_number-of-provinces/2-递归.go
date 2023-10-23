package main

func main() {

}

var arr []bool

func findCircleNum(M [][]int) int {
	n := len(M)
	arr = make([]bool, n)
	res := 0
	for i := 0; i < n; i++ {
		if arr[i] == false {
			dfs(M, i)
			res++
		}
	}
	return res
}

func dfs(M [][]int, index int) {
	for i := 0; i < len(M); i++ {
		if arr[i] == false && M[index][i] == 1 {
			arr[i] = true
			dfs(M, i)
		}
	}
}
