package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bulbSwitch(12))
}

// leetcode319_灯泡开关
func bulbSwitch(n int) int {
	// 第i个灯泡的反转次数等于它所有因子（包括1和i）的个数
	// 反转奇数次=>变成亮
	// 只有平方数才有奇数个因子
	return int(math.Sqrt(float64(n)))
}
