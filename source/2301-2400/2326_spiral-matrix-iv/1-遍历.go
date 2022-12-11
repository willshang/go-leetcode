package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 顺时针：上右下左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = -1
		}
	}

	x, y, dir := 0, 0, 0
	for ; head != nil; head = head.Next {
		res[x][y] = head.Val
		newX, newY := x+dx[dir], y+dy[dir]
		if 0 > newX || newX >= m || 0 > newY || newY >= n || res[newX][newY] != -1 {
			dir = (dir + 1) % 4 // 换方向
		}
		x = x + dx[dir]
		y = y + dy[dir]
	}
	return res
}
