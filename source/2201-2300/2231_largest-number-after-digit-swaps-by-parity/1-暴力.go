package main

import (
	"strconv"
)

func main() {

}

// leetcode2231_按奇偶性交换后的最大数字
func largestInteger(num int) int {
	value := []byte(strconv.Itoa(num))
	n := len(value)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if int(value[i]-'0')%2 == int(value[j]-'0')%2 &&
				value[i] < value[j] {
				value[i], value[j] = value[j], value[i]
			}
		}
	}
	res, _ := strconv.Atoi(string(value))
	return res
}
