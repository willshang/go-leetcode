package main

import "fmt"

func main() {
	//fmt.Println(drawLine(3,96,0,95,0))
	fmt.Println(drawLine(15, 96, 81, 95, 1))
}

func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	arr := make([]int32, length)
	width := w / 32
	for i := x1; i <= x2; i++ {
		index := width*y + (i / 32)
		arr[index] = arr[index] ^ (1 << (31 - (i % 32)))
	}
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = int(arr[i])
	}
	return res
}
