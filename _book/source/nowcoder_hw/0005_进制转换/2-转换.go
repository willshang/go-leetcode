package main

import "fmt"

func main() {
	var num int
	for {
		n, _ := fmt.Scanf("%v", &num)
		if n == 0 {
			break
		}
		fmt.Println(num)
	}
}
