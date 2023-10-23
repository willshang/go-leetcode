package main

func main() {

}

// leetcode1222_可以攻击国王的皇后
var dir = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	arr := make([][]int, 8)
	for i := 0; i < 8; i++ {
		arr[i] = make([]int, 8)
	}
	for i := 0; i < len(queens); i++ {
		a, b := queens[i][0], queens[i][1]
		arr[a][b] = 1
	}
	res := make([][]int, 0)
	for i := 0; i < len(dir); i++ {
		x := dir[i][0] + king[0]
		y := dir[i][1] + king[1]
		for 0 <= x && x < 8 && 0 <= y && y < 8 {
			if arr[x][y] == 1 {
				res = append(res, []int{x, y})
				break
			}
			x = x + dir[i][0]
			y = y + dir[i][1]
		}
	}
	return res
}
