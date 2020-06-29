package main

import "fmt"

func main() {
	var str string
	for {
		flag, _ := fmt.Scan(&str)
		if flag == 0 {
			break
		}
	}
	fmt.Println(len(str))
}
