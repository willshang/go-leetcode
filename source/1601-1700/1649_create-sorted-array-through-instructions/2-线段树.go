package main

import "fmt"

func main() {
	fmt.Println(createSortedArray([]int{1, 5, 6, 2}))
}

var mod = 1000000007

func createSortedArray(instructions []int) int {
	res := 0
	n := 100001
	arr = make([]int, n*4+1)
	for i := 0; i < len(instructions); i++ {
		x := instructions[i]
		left := query(0, 1, n, 1, x-1)  // 查询1~x-1
		right := query(0, 1, n, x+1, n) // 查询 x+1~n
		res = (res + min(left, right)) % mod
		update(0, 1, n, x) // 添加x
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var arr []int // 线段树

func update(id int, left, right, x int) {
	if left > x || right < x {
		return
	}
	arr[id]++
	if left == right {
		return
	}
	mid := left + (right-left)/2
	update(2*id+1, left, mid, x)    // 左节点
	update(2*id+2, mid+1, right, x) // 右节点
}

func query(id int, left, right, queryLeft, queryRight int) int {
	if left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		return arr[id]
	}
	mid := left + (right-left)/2
	return query(id*2+1, left, mid, queryLeft, queryRight) +
		query(id*2+2, mid+1, right, queryLeft, queryRight)
}
