package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 72, 75, 71, 69, 72, 76, 73}))
}

// 剑指OfferII038.每日温度
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0) // 栈保存递减数据的下标
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[last] = i - last
		}
		stack = append(stack, i)
	}
	return res
}
