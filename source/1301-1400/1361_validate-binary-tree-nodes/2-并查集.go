package main

import "fmt"

func main() {
	fmt.Println(validateBinaryTreeNodes(4, []int{3, -1, 1 - 1}, []int{-1, -1, 0, -1}))
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	fa = Init(n)
	for i := 0; i < n; i++ {
		a, b := leftChild[i], rightChild[i]
		if a != -1 {
			if find(a) != a || query(i, a) == true {
				return false
			}
			union(a, i) // 注意顺序
		}
		if b != -1 {
			if find(b) != b || query(i, b) == true {
				return false
			}
			union(b, i) // 注意顺序
		}
	}
	return getCount() == 1
}

var fa []int
var count int

// 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	count = n
	return arr
}

// 查询
func find(x int) int {
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}

// 合并
func union(i, j int) {
	x, y := find(i), find(j)
	if x != y {
		fa[x] = y
		count--
	}
}

func query(i, j int) bool {
	return find(i) == find(j)
}

func getCount() int {
	return count
}
