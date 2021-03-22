package main

import "fmt"

func main() {
	fmt.Println(numSteps("1101"))
}

// leetcode1404_将二进制表示减到1的步骤数
func numSteps(s string) int {
	res := 0
	for s != "1" {
		n := len(s)
		if s[n-1] == '0' {
			s = s[:n-1]
		} else {
			s = add(s)
		}
		res++
	}
	return res
}

func add(s string) string {
	arr := []byte(s)
	flag := true
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == '0' {
			arr[i] = '1'
			flag = false
		} else {
			arr[i] = '0'
		}
		if flag == false {
			return string(arr)
		}
	}
	return "1" + string(arr)
}
