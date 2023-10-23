package main

import "fmt"

func main() {
	//fmt.Println(validUtf8([]int{235, 140, 4}))
	//fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(validUtf8([]int{145}))
	//fmt.Println(validUtf8([]int{248, 130, 130, 130}))
}

// leetcode393_UTF-8编码验证
func validUtf8(data []int) bool {
	count := 0
	for i := 0; i < len(data); i++ {
		if count == 0 {
			if data[i]&0b11111000 == 0b11110000 {
				count = 3
			} else if data[i]&0b11110000 == 0b11100000 {
				count = 2
			} else if data[i]&0b11100000 == 0b11000000 {
				count = 1
			} else if data[i]&0b10000000 != 0b00000000 { // 其它以1开头的
				return false
			}
		} else {
			if data[i]&0b10000000 != 0b10000000 {
				return false
			}
			count--
		}
	}
	return count == 0
}
