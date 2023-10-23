package main

import "strings"

func main() {

}

// leetcode1324_竖直打印单词
func printVertically(s string) []string {
	arr := strings.Split(s, " ")
	n := len(arr)
	maxLength := 0
	for i := 0; i < n; i++ {
		if len(arr[i]) > maxLength {
			maxLength = len(arr[i])
		}
	}
	temp := make([][]byte, maxLength)
	for i := 0; i < maxLength; i++ {
		temp[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			temp[i][j] = ' '
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(arr[i]); j++ {
			temp[j][i] = arr[i][j]
		}
	}
	res := make([]string, 0)
	for i := 0; i < len(temp); i++ {
		res = append(res, strings.TrimRight(string(temp[i]), " "))
	}
	return res
}
