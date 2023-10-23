package main

import "sort"

func main() {

}

func minSessions(tasks []int, sessionTime int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	left, right := 1, len(tasks)

	for left < right {
		mid := left + (right-left)/2
		arr := make([]int, mid)
		if dfs(tasks, sessionTime, 0, mid, arr) == true {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func dfs(tasks []int, sessionTime int, index, count int, arr []int) bool {
	if index == len(tasks) { // 到最后退出
		return true
	}
	flag := false
	for i := 0; i < count; i++ { // 遍历每个工作段
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
		if dfs(tasks, sessionTime, index+1, count, arr) == true {
			return true
		}
		arr[i] = arr[i] - tasks[index]
	}
	return false
}
