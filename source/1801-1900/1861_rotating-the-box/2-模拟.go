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

func rotateTheBox(box [][]byte) [][]byte {
	n, m := len(box), len(box[0])
	for i := 0; i < n; i++ {
		queue := make([]int, 0)
		for j := m - 1; j >= 0; j-- {
			if box[i][j] == '*' {
				queue = make([]int, 0) // 维护可移动的y坐标
			} else if box[i][j] == '#' {
				if len(queue) > 0 { // 石头落下
					first := queue[0]        // 最右边能落下的位置
					queue = queue[1:]        // 退出队列
					box[i][first] = '#'      // 落下
					box[i][j] = '.'          // 原位置为空
					queue = append(queue, j) // 置空后进入队列
				}
			} else {
				queue = append(queue, j)
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
