package main

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func longestIncreasingPath(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	queue := make([][2]int, 0) // 从最大数开始广度优先搜索
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				x, y := i+dx[k], j+dy[k]
				if 0 <= x && x < n && 0 <= y && y < m && matrix[x][y] > matrix[i][j] {
					arr[i][j]++ // 四周大于当前的个数
				}
			}
			if arr[i][j] == 0 { // 四周没有大于当前的数
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	res := 0
	for len(queue) > 0 {
		res++
		length := len(queue)
		for i := 0; i < length; i++ {
			a, b := queue[i][0], queue[i][1]
			for k := 0; k < 4; k++ {
				x, y := a+dx[k], b+dy[k]
				if 0 <= x && x < n && 0 <= y && y < m && matrix[a][b] > matrix[x][y] {
					arr[x][y]--
					if arr[x][y] == 0 { // 个数为0，加入队列
						queue = append(queue, [2]int{x, y})
					}
				}
			}
		}
		queue = queue[length:]
	}
	return res
}
