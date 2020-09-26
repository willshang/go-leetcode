package main

import "fmt"

func main() {
	fmt.Println(swapNumbers([]int{1, 2}))
}

// 程序员面试金典16.01_交换数字
func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] + numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] = numbers[0] - numbers[1]
	return numbers
}
