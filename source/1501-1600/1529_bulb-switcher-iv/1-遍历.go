package main

import "fmt"

func main() {
	fmt.Println(minFlips("10111"))
}

func minFlips(target string) int {
	res := 0
	prev := uint8('0')
	for i := 0; i < len(target); i++ {
		// 只要后一个与前一个不相等, 都会额外增加一次翻转
		if prev != target[i] {
			res++
			prev = target[i]
		}
	}
	return res
}
