package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(minimizeResult("247+38"))
}

// leetcode2232_向表达式添加括号后的最小结果
func minimizeResult(expression string) string {
	res := ""
	arr := strings.Split(expression, "+")
	left, right := arr[0], arr[1]
	minValue := math.MaxInt32
	for i := 0; i < len(left); i++ {
		for j := 1; j <= len(right); j++ {
			var a, b, c, d int
			if i == 0 { // 左边
				a = 1
			} else {
				a, _ = strconv.Atoi(left[:i])
			}
			b, _ = strconv.Atoi(left[i:])
			c, _ = strconv.Atoi(right[:j])
			d = 1
			if j < len(right) { // 右边
				d, _ = strconv.Atoi(right[j:])
			}
			if a*(b+c)*d < minValue {
				minValue = a * (b + c) * d
				res = fmt.Sprintf("%s(%s+%s)%s", left[:i], left[i:], right[:j], right[j:])
			}
		}
	}
	return res
}
