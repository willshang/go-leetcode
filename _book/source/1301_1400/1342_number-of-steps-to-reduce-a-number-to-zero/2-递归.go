package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
}

var res int

func numberOfSteps(num int) int {
	res = 0
	dfs(num)
	return res
}

func dfs(num int) {
	if num != 0 {
		res++
		if num%2 == 1 {
			dfs(num - 1)
		} else {
			dfs(num / 2)
		}
	}
}
