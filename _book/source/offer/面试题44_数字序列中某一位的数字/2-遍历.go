package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(12))
}

// 剑指offer_面试题44_数字序列中某一位的数字
/*
1-9 	9*1 1位
10-99 	90*2 2位
100-999 900*3 3位
*/
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	digits := 1
	count := 9
	number := 1
	for n-digits*count > 0 {
		n = n - digits*count
		digits++
		count = count * 10
		number = number * 10
	}
	number = number + (n-1)/digits
	index := (n - 1) % digits
	str := strconv.Itoa(number)
	return int(str[index] - '0')
}
