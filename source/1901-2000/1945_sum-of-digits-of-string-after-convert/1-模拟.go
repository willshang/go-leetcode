package main

import "fmt"

func main() {
	fmt.Println(getLucky("iiiiz", 1))
}

func getLucky(s string, k int) int {
	arr := make([]int, 0)
	for i := 0; i < len(s); i++ {
		value := int(s[i]-'a') + 1
		if value < 10 {
			arr = append(arr, value)
		} else {
			arr = append(arr, value/10)
			arr = append(arr, value%10)
		}
	}
	for i := 1; i <= k; i++ {
		arr = trans(arr)
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		res = 10*res + arr[i]
	}
	return res
}

func trans(arr []int) []int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	res := make([]int, 0)
	for sum > 0 {
		res = append([]int{sum % 10}, res...)
		sum = sum / 10
	}
	return res
}
