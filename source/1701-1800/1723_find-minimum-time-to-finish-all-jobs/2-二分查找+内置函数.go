package main

import "sort"

func main() {

}

func minimumTimeRequired(jobs []int, k int) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] > jobs[j]
	})
	sum := 0
	for i := 0; i < len(jobs); i++ {
		sum = sum + jobs[i]
	}
	left, right := jobs[0], sum
	return left + sort.Search(right-left, func(limit int) bool {
		return dfs(jobs, make([]int, k), limit+left, 0)
	})
}

func dfs(jobs []int, arr []int, limit int, index int) bool {
	if index >= len(jobs) {
		return true
	}
	for i := 0; i < len(arr); i++ {
		if arr[i]+jobs[index] <= limit {
			arr[i] = arr[i] + jobs[index]
			if dfs(jobs, arr, limit, index+1) == true {
				return true
			}
			arr[i] = arr[i] - jobs[index]
		}
		// 当前没有被分配 | 当前分配达到上线 => 不需要尝试继续分配
		// 剪枝:去除也可以过
		if arr[i] == 0 || arr[i]+jobs[index] == limit {
			break
		}
	}
	return false
}
