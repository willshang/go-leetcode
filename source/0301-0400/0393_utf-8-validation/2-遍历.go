package main

import "fmt"

func main() {
	//fmt.Println(validUtf8([]int{235, 140, 4}))
	//fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(validUtf8([]int{145}))
	//fmt.Println(validUtf8([]int{248, 130, 130, 130}))
}

func validUtf8(data []int) bool {
	count := 0
	for i := 0; i < len(data); i++ {
		if count == 0 {
			if data[i]>>3 == 0b11110 {
				count = 3
			} else if data[i]>>4 == 0b1110 {
				count = 2
			} else if data[i]>>5 == 0b110 {
				count = 1
			} else if data[i]>>7 != 0b0 { // 其它以1开头的
				return false
			}
		} else {
			if data[i]>>6 != 0b10 {
				return false
			}
			count--
		}
	}
	return count == 0
}
