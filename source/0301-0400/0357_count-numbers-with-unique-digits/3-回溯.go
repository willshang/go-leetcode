package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(10))
	fmt.Println(countNumbersWithUniqueDigits(2))
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	return dfs(n, 0, make([]bool, 10))
}

func dfs(n, index int, visited []bool) int {
	if index == n {
		return 0
	}
	res := 0
	for i := 0; i < 10; i++ {
		if n >= 2 && index == 1 && i == 0 {
			continue
		}
		if visited[i] == true {
			continue
		}
		visited[i] = true
		res = res + dfs(n, index+1, visited) + 1
		visited[i] = false
	}
	return res
}
