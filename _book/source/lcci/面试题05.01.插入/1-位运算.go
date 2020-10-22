package main

import "fmt"

func main() {
	fmt.Println(insertBits(0, 31, 0, 4))
	//fmt.Println(insertBits(1024, 19, 2, 6))
}

// 程序员面试金典05.01_插入
func insertBits(N int, M int, i int, j int) int {
	a := (N >> (j + 1)) << (j + 1)
	b := (N>>i)<<i ^ N
	c := M << i
	return a | b | c
}
