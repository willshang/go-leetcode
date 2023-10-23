package main

func main() {

}

// leetcode1776_车队II
func getCollisionTimes(cars [][]int) []float64 {
	n := len(cars)
	res := make([]float64, n)
	stack := make([]int, 0) // 单调栈
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			a := float64(cars[top][0] - cars[i][0]) // 距离差
			b := float64(cars[i][1] - cars[top][1]) // 速度差
			if (cars[i][1] <= cars[top][1]) ||      // 1、永远追不上栈顶车
				(res[top] > 0 && a/b > res[top]) { // 2、可以在栈顶车追上更前面车之前追上栈顶车
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		if len(stack) == 0 {
			res[i] = -1
		} else { // 可以追上
			top := stack[len(stack)-1]
			a := float64(cars[top][0] - cars[i][0]) // 距离差
			b := float64(cars[i][1] - cars[top][1]) // 速度差
			res[i] = a / b
		}
		stack = append(stack, i)
	}
	return res
}
