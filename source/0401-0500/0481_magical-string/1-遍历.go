package main

import "fmt"

func main() {
	fmt.Println(magicalString(20))
}

func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}
	str := []byte("122")
	res := 1
	index := 2
	for i := 2; i < n; i++ {
		if str[i] == '2' {
			if str[index] == '2' {
				str = append(str, []byte{'1', '1'}...)
			} else {
				str = append(str, []byte{'2', '2'}...)
			}
			index = index + 2
		} else {
			res++
			if str[index] == '2' {
				str = append(str, '1')
			} else {
				str = append(str, '2')
			}
			index = index + 1
		}
	}
	return res
}
