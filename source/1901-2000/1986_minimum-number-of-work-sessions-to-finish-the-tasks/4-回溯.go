package main

import "fmt"

func main() {
	a := 21
	for i := a; i > 0; i = (i - 1) & a {
		fmt.Printf("%b\n", i)
	}
}

var res int

func minSessions(tasks []int, sessionTime int) int {
	res = len(tasks)
	dfs(tasks, sessionTime, 0, 0, make([]int, len(tasks)))
	return res
}

func dfs(tasks []int, sessionTime int, index, count int, arr []int) {
	if count >= res {
		return
	}
	if index == len(tasks) {
		res = count
		return
	}
	flag := false
	for i := 0; i < len(tasks); i++ { // 尝试每个工作段
		if arr[i] == 0 && flag == true { // 使用1个新的工作段即可
			break
		}
		if arr[i]+tasks[index] > sessionTime { // 当前超时，跳过尝试新的工作段
			continue
		}
		if arr[i] == 0 {
			flag = true
		}
		arr[i] = arr[i] + tasks[index]
		if flag == true {
			dfs(tasks, sessionTime, index+1, count+1, arr) // 有使用新的工作段
		} else {
			dfs(tasks, sessionTime, index+1, count, arr) // 没有使用新的工作段
		}
		arr[i] = arr[i] - tasks[index]
	}
}
