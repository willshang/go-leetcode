package main

import "fmt"

func main() {
	fmt.Println(magicalString(20))
}

// leetcode481_神奇字符串
func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}
	str := []byte("122")
	flag := true
	for i := 2; i < n; i++ {
		count := str[i] - '0'
		if flag == true {
			for count > 0 {
				str = append(str, '1')
				count--
			}
			flag = false
		} else {
			for count > 0 {
				str = append(str, '2')
				count--
			}
			flag = true
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if str[i] == '1' {
			res++
		}
	}
	return res
}
