package main

import (
	"strconv"
	"strings"
)

func main() {

}

// leetcode1432_改变一个整数能得到的最大差值
func maxDiff(num int) int {
	maxValue, minValue := num, num
	str := strconv.Itoa(num)
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			newStr := strings.ReplaceAll(str, string('0'+x), string('0'+y))
			if newStr[0] == '0' {
				continue
			}
			value, _ := strconv.Atoi(newStr)
			if value > maxValue {
				maxValue = value
			}
			if value < minValue {
				minValue = value
			}
		}
	}
	return maxValue - minValue
}
