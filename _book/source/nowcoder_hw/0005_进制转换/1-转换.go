package main

import "fmt"

func main() {
	var num int
	for {
		n, _ := fmt.Scanf("0x%x", &num)
		if n == 0 {
			break
		}
		fmt.Println(num)
	}
}
