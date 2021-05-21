package main

func main() {

}

// leetcode909_蛇梯棋
func snakesAndLadders(board [][]int) int {
	n := len(board)
	m := make(map[int]int) // 保存到达坐标对应的移动次数
	m[1] = 0
	queue := make([]int, 0)
	queue = append(queue, 1)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == n*n {
			return m[node]
		}
		for i := node + 1; i <= node+6 && i <= n*n; i++ {
			_, a, b := getIndex(i, n)
			next := -1
			if board[a][b] == -1 {
				next = i
			} else {
				next = board[a][b]
			}
			if _, ok := m[next]; !ok {
				m[next] = m[node] + 1
				queue = append(queue, next)
			}
		}
	}
	return -1
}

func getIndex(i int, n int) (int, int, int) {
	var x, y int
	x = (i - 1) / n // 行
	y = (i - 1) % n // 列
	if x%2 == 1 {   // 奇数行需要反转
		y = n - 1 - y // 反转
	}
	x = n - 1 - x // 反转
	return x*n + y, x, y
}
