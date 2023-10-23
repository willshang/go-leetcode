package main

import "fmt"

func main() {
	fmt.Println(multiply("123", "456"))
}

// leetcode43_字符串相乘
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	arr := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		a := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			b := int(num2[j] - '0')
			value := a*b + arr[i+j+1]
			arr[i+j+1] = value % 10
			arr[i+j] = value/10 + arr[i+j]
		}
	}
	res := ""
	for i := 0; i < len(arr); i++ {
		if i == 0 && arr[i] == 0 {
			continue
		}
		res = res + string(arr[i]+'0')
	}
	return res
}
