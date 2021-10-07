package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int)
	res := 1
	fa = Init(nums)
	for i := 0; i < len(nums); i++ {
		union(nums[i], nums[i]+1)
		m[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		res = max(res, find(nums[i])-nums[i]+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var fa map[int]int

// 初始化
func Init(data []int) map[int]int {
	n := len(data)
	arr := make(map[int]int)
	for i := 0; i < n; i++ {
		arr[data[i]] = data[i]
	}
	return arr
}

// 查询
func find(x int) int {
	if _, ok := fa[x]; !ok {
		return math.MinInt32 // 特殊处理
	}
	res := x
	for res != fa[res] {
		res = fa[res]
	}
	return res
}

// 合并
func union(i, j int) {
	x, y := find(i), find(j)
	if x == y {
		return
	} else if x == math.MinInt32 || y == math.MinInt32 {
		return
	}
	fa[x] = y
}

func query(i, j int) bool {
	return find(i) == find(j)
}
