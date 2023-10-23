package main

import "fmt"

func main() {
	fmt.Println(minFlips("10111"))
}

// leetcode1529_灯泡开关IV
func minFlips(target string) int {
	res := 0
	target = "0" + target
	for i := 1; i < len(target); i++ {
		// 只要后一个与前一个不相等, 都会额外增加一次翻转
		// 从前往后，每次翻转一样的
		// 如: "10111" => 00000=>(1)1111=>(10)000=>(10111)
		if target[i] != target[i-1] {
			res++
		}
	}
	return res
}
