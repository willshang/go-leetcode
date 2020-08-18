package main

import "fmt"

func main() {
	fmt.Println(printBin(0.625))
}

// 程序员面试金典05.02_二进制数转字符串
func printBin(num float64) string {
	res := "0."
	for num != float64(0) {
		num = num * 2
		if num >= 1 {
			res = res + "1"
			num = num - 1.0
		} else {
			res = res + "0"
		}
		if len(res) > 32 {
			return "ERROR"
		}
	}
	return res
}
