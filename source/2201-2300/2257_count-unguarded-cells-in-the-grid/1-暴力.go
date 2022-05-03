package main

func main() {

}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < len(walls); i++ {
		a, b := walls[i][0], walls[i][1]
		arr[a][b] = 'W'
	}
	for i := 0; i < len(guards); i++ {
		a, b := guards[i][0], guards[i][1]
		arr[a][b] = 'G'
	}
	for i := 0; i < len(guards); i++ {
		a, b := guards[i][0], guards[i][1]
		x, y := a, b-1
		for y >= 0 && arr[x][y] != 'G' && arr[x][y] != 'W' {
			arr[x][y] = 'B'
			y--
		}
		x, y = a-1, b
		for x >= 0 && arr[x][y] != 'G' && arr[x][y] != 'W' {
			arr[x][y] = 'B'
			x--
		}
		x, y = a, b+1
		for y < n && arr[x][y] != 'G' && arr[x][y] != 'W' {
			arr[x][y] = 'B'
			y++
		}
		x, y = a+1, b
		for x < m && arr[x][y] != 'G' && arr[x][y] != 'W' {
			arr[x][y] = 'B'
			x++
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				res++
			}
		}
	}
	return res
}
