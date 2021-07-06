package main

import "fmt"

func main() {
	//fmt.Println(drawLine(3,96,0,95,0))
	fmt.Println(drawLine(15, 96, 81, 95, 1))
}

// 程序员面试金典05.08_绘制直线
func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		start := i * 32
		var value int32
		for j := 0; j < 32; j++ {
			if x1+y*w <= start+j && start+j <= x2+y*w {
				value = value ^ (1 << (31 - j)) // 画线：把第31-j位变为1
			}
		}
		res[i] = int(value)
	}
	return res
}
