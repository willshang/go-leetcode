package main

import "fmt"

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}

func arrayNesting(nums []int) int {
	res := 0
	fa = Init(len(nums))
	for i := 0; i < len(nums); i++ {
		union(i, nums[i])
	}
	m := make(map[int]int)
	for i := 0; i < len(fa); i++ {
		m[find(i)]++
	}
	for _, v := range m {
		if v > res {
			res = v
		}
	}
	return res
}

var fa []int

// 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// 查询
func find(x int) int {
	if fa[x] == x {
		return x
	}
	// 路径压缩
	fa[x] = find(fa[x])
	return fa[x]
}

// 合并
func union(i, j int) {
	x, y := find(i), find(j)
	if x != y {
		fa[x] = y
	}
}
