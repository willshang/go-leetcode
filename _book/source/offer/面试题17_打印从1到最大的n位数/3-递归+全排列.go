package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(printNumbers(1))
}

var temp []string

func printNumbers(n int) []int {
	if n <= 0 {
		return nil
	}
	temp = make([]string, 0)
	arr := make([]rune, n)
	for i := 0; i < 10; i++ {
		arr[0] = rune(i + '0')
		dfs(arr, n, 0)
	}

	res := make([]int, 0)
	for i := 0; i < len(temp); i++ {
		value, _ := strconv.Atoi(temp[i])
		res = append(res, value)
	}
	return res
}

func dfs(arr []rune, n int, index int) {
	if index == n-1 {
		if printNum(arr) != "" {
			temp = append(temp, printNum(arr))
		}
		return
	}
	for i := 0; i < 10; i++ {
		arr[index+1] = rune(i + '0')
		dfs(arr, n, index+1)
	}
}

func printNum(arr []rune) string {
	res := ""
	isBeginning := true
	for i := 0; i < len(arr); i++ {
		if isBeginning && arr[i] != '0' {
			isBeginning = false
		}
		if !isBeginning {
			res = res + string(arr[i])
		}
	}
	return res
}
