package main

import "strings"

func main() {

}

// leetcode2018_判断单词是否能放入填字游戏内
func placeWordInCrossword(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	arr := make([][]byte, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[j][i] = board[i][j] // 转置后统一按行处理
		}
	}
	return checkBoard(board, word) || checkBoard(arr, word)
}

func checkBoard(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		arr := strings.Split(string(board[i]), "#") // 按#切割
		for j := 0; j < len(arr); j++ {
			if len(arr[j]) != len(word) { // 长度不等
				continue
			}
			left, right := true, true
			for k := 0; k < len(word); k++ { // 正向：word == arr[j]
				if arr[j][k] != ' ' && arr[j][k] != word[k] {
					left = false
					break
				}
			}
			for k := 0; k < len(word); k++ { // 反向：word == reverse(arr[j])
				if arr[j][k] != ' ' && arr[j][k] != word[len(word)-1-k] {
					right = false
					break
				}
			}
			if left == true || right == true {
				return true
			}
		}
	}
	return false
}
