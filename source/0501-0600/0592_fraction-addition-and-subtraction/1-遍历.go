package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

// leetcode592_分数加减运算
func fractionAddition(expression string) string {
	s := strings.ReplaceAll(expression, "-", "+-") // 替换-为+-
	s = strings.ReplaceAll(s, "/", "+")            // 替换/为+
	temp := strings.Split(s, "+")                  // 根据+切割
	arr := make([]int, 0)
	for i := 0; i < len(temp); i++ {
		if temp[i] == "" { // 第一个数的分子如果为负数，+切割会为空
			continue
		}
		value, _ := strconv.Atoi(temp[i])
		arr = append(arr, value)
	}
	a, b := 0, 1
	for i := 0; i < len(arr); i = i + 2 { // 遍历每个数：分子+分母
		c, d := arr[i], arr[i+1]
		// a/b+c/d=ad/bd+bc/bd=(ad+bc)/bd
		a = a*d + b*c
		b = b * d
		g := gcd(a, b)
		if g < 0 { // 约分为负数，需要转换为正数，这样确保分母一直为正数
			g = -g
		}
		a, b = a/g, b/g
	}
	return fmt.Sprintf("%d/%d", a, b)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
