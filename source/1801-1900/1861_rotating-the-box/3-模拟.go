package main

import "fmt"

func main() {
	arr := [][]byte{
		{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '*', '#', '*'},
	}
	fmt.Println(rotateTheBox(arr))
}

// leetcode1861_旋转盒子
func rotateTheBox(box [][]byte) [][]byte {
	n, m := len(box), len(box[0])
	for i := 0; i < n; i++ {
		last := m - 1 // 最后1个空位
		for j := m - 1; j >= 0; j-- {
			if box[i][j] == '*' {
				last = j - 1 // 有障碍，最后空位更新
			} else if box[i][j] == '#' {
				if last > j { // 可以移动
					box[i][last] = '#'
					box[i][j] = '.'
					last--
				} else { // 当前位置不可以移动
					last = j - 1
				}
			}
		}
	}
	res := make([][]byte, m)
	for i := 0; i < m; i++ {
		res[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[j][n-1-i] = box[i][j]
		}
	}
	return res
}
