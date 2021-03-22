package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	arr := make([]int, 0)
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10
	}
	res := 0
	flag := true
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 6 && flag == true {
			res = res*10 + 9
			flag = false
		} else {
			res = res*10 + arr[i]
		}
	}
	return res
}
