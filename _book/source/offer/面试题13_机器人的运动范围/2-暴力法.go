package main

import "fmt"

func main() {
	fmt.Println(movingCount(1, 2, 1))
}

// 剑指offer_面试题13_机器人的运动范围
func movingCount(m int, n int, k int) int {
	if k < 0 || m <= 0 || n <= 0 {
		return 0
	}
	res := 1
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i-1 >= 0 && visited[i-1][j] == true) || (j-1 >= 0 && visited[i][j-1] == true) {
				value := getDigiSum(i) + getDigiSum(j)
				if value <= k {
					res++
					visited[i][j] = true
				}
			}
		}
	}
	return res
}

func getDigiSum(num int) int {
	sum := 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}
