package main

import (
	"fmt"
)

func main() {
	var a, b, s int
	fmt.Scan(&a, &b, &s)
	res := getResult(a, b, s)
	if res == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func getResult(a, b, s int) bool {
	a = abs(a)
	b = abs(b)
	if a+b <= s && (s-a-b)%2 == 0 { // 注意2个条件
		return true
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
