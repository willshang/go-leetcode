package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(arr))
}

// leetcode498_对角线遍历
func findDiagonalOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	n, m := len(matrix), len(matrix[0])
	if n == 1 {
		return matrix[0]
	}
	i, j := 0, 0
	flag := false
	for j < m {
		a, b := i, j
		temp := make([]int, 0)
		temp = append(temp, matrix[a][b])
		// 从左下往右上
		for a != 0 && b != m-1 {
			a--
			b++
			temp = append(temp, matrix[a][b])
		}
		if flag == true {
			reverse(temp)
			flag = false
		} else {
			flag = true
		}
		res = append(res, temp...)
		if i != n-1 {
			i++
		} else {
			j++
		}
	}
	return res
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
