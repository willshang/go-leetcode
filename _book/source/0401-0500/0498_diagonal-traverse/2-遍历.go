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

func findDiagonalOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	n, m := len(matrix), len(matrix[0])
	if n == 1 {
		return matrix[0]
	}
	// 右边拼接一个相同的，依次遍历
	for i := 0; i < n+m-1; i++ {
		temp := make([]int, 0)
		var a, b int
		if i < m {
			a = 0
			b = i
		} else {
			a = i - m + 1
			b = m - 1
		}
		for a < n && b >= 0 {
			temp = append(temp, matrix[a][b])
			a, b = a+1, b-1
		}
		if i%2 == 0 {
			reverse(temp)
		}
		res = append(res, temp...)
	}
	return res
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
