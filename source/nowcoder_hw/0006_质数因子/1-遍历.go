package main

import "fmt"

func main() {
	var num int64
	for {
		n, _ := fmt.Scanf("%d", &num)
		if n == 0 {
			break
		}
		i := int64(2)
		for num > 1 {
			if num%i == 0 {
				fmt.Printf("%d ", i)
				num = num / i
			} else {
				i++
			}
		}
		fmt.Println()
	}
}
