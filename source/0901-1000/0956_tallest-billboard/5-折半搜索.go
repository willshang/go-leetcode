package main

import "fmt"

func main() {
	fmt.Println(tallestBillboard([]int{1, 3, 5, 7}))
}

func tallestBillboard(rods []int) int {
	leftM := makeM(rods, 0, len(rods)/2)
	rightM := makeM(rods, len(rods)/2, len(rods))
	res := 0
	for k := range leftM {
		if rightM[k] > 0 {
			res = max(res, leftM[k]+rightM[-k])
		}
	}
	return res
}

func makeM(rods []int, left, right int) map[int]int {
	dp := make([][2]int, 100001)
	dp[0] = [2]int{0, 0}
	count := 1
	for i := left; i < right; i++ {
		length := count
		for j := 0; j < length; j++ {
			dp[count] = [2]int{dp[j][0] + rods[i], dp[j][1]}
			count++
			dp[count] = [2]int{dp[j][0], dp[j][1] + rods[i]}
			count++
		}
	}
	m := make(map[int]int)
	for i := 0; i < count; i++ {
		a := dp[i][0]
		b := dp[i][1]
		m[a-b] = max(m[a-b], a)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
