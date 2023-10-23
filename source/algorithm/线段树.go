package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	n := len(nums)
	arr = make([]int, 4*n+1)
	for i := 0; i < n; i++ {
		update(0, 1, n, nums[i], 1) //  添加nums[i]
	}
	res := query(0, 1, n, 1, 3) // 查询1~3
	fmt.Println(res)
}

var arr []int // 线段树

func update(id int, left, right, x int, value int) {
	if left > x || right < x {
		return
	}
	arr[id] = arr[id] + value
	if left == right {
		return
	}
	mid := left + (right-left)/2
	update(2*id+1, left, mid, x, value)    // 左节点
	update(2*id+2, mid+1, right, x, value) // 右节点
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
