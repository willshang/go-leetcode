package main

import "fmt"

func main() {
	fmt.Println(movingCount(2, 3, 1))
}

func movingCount(m int, n int, k int) int {
	if k < 0 || m <= 0 || n <= 0 {
		return 0
	}
	visited := make([]bool, m*n)
	res := dfs(m, n, k, visited, 0, 0)
	return res
}

func dfs(m, n, k int, visited []bool, x, y int) int {
	res := 0
	if check(m, n, k, visited, x, y) {
		visited[x*n+y] = true
		res = 1 + dfs(m, n, k, visited, x+1, y) +
			dfs(m, n, k, visited, x-1, y) +
			dfs(m, n, k, visited, x, y+1) +
			dfs(m, n, k, visited, x, y-1)
	}
	return res
}

func check(m, n, k int, visited []bool, x, y int) bool {
	if x >= 0 && x < m && y >= 0 && y < n &&
		getDigiSum(x)+getDigiSum(y) <= k && visited[x*n+y] == false {
		return true
	}
	return false
}

func getDigiSum(num int) int {
	sum := 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}
