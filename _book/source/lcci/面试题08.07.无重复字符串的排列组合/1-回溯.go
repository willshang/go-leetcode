package main

import "fmt"

func main() {
	fmt.Println(permutation("qwe"))
}

var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	visited := make(map[int]bool)
	dfs(nums, 0, "", visited)
	return res
}

func dfs(nums []byte, index int, str string, visited map[int]bool) {
	if index == len(nums) {
		res = append(res, str)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == false {
			str = str + string(nums[i])
			visited[i] = true
			dfs(nums, index+1, str, visited)
			str = str[:len(str)-1]
			visited[i] = false
		}
	}
}
