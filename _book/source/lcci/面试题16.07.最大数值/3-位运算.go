package main

import (
	"fmt"
)

func main() {
	//fmt.Println(maximum(1, 2))
	fmt.Println(maximum(-777777, -2222))
}

// 程序员面试金典16.07_最大数值
func maximum(a int, b int) int {
	value := int(uint64(a-b) >> 63) // 取符号位，a-b>0 => 符号位为0 a-b<0 =>符号位为1
	return value*b + int(1^value)*a // value=0=> 0^1=1 1^1=0
}
