package main

import "fmt"

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}

// leetcode621_任务调度器
func leastInterval(tasks []byte, n int) int {
	arr := [26]int{}
	maxValue := 0
	for i := 0; i < len(tasks); i++ {
		arr[tasks[i]-'A']++
		if arr[tasks[i]-'A'] > maxValue {
			maxValue = arr[tasks[i]-'A']
		}
	}
	res := (maxValue - 1) * (n + 1) // 完成所有任务至少需要(max-1)*(n+1)+1
	for i := 0; i < len(arr); i++ {
		if arr[i] == maxValue {
			res++
		}
	}
	return max(res, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
