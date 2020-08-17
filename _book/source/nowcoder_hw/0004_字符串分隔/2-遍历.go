package main

import (
	"fmt"
)

func main() {
	arr := [2]string{}
	fmt.Scanln(&arr[0])
	fmt.Scanln(&arr[1])
	for i := 0; i < 2; i++ {
		str := arr[i]
		length := len(str)
		if length%8 != 0 {
			for i := 0; i < 8-length%8; i++ {
				str = str + string('0')
			}
		}
		for i := 0; i < len(str)/8; i++ {
			fmt.Println(str[i*8 : i*8+8])
		}
	}
}
