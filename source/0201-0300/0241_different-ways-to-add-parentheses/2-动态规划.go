package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
}

func diffWaysToCompute(input string) []int {
	if value, err := strconv.Atoi(input); err == nil {
		return []int{value}
	}
	numArr := make([]int, 0)
	opArr := make([]byte, 0)
	num := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			opArr = append(opArr, input[i])
			numArr = append(numArr, num)
			num = 0
			continue
		}
		num = num*10 + int(input[i]-'0')
	}
	numArr = append(numArr, num)
	n := len(numArr)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		arr := make([]int, 0)
		arr = append(arr, numArr[i])
		dp[i][i] = arr
	}
	for k := 2; k <= n; k++ { // 长度
		for i := 0; i < n; i++ { // 起点
			j := i + k - 1 // 终点
			if j >= n {
				break
			}
			temp := make([]int, 0)
			for l := i; l < j; l++ { // 切割点
				left := dp[i][l]
				right := dp[l+1][j]
				for a := 0; a < len(left); a++ {
					for b := 0; b < len(right); b++ {
						op := opArr[l]
						if op == '+' {
							temp = append(temp, left[a]+right[b])
						} else if op == '-' {
							temp = append(temp, left[a]-right[b])
						} else if op == '*' {
							temp = append(temp, left[a]*right[b])
						}
					}
				}
			}
			dp[i][j] = temp
		}
	}
	return dp[0][n-1]
}
