package main

import "fmt"

func main() {
	var num int64
	for {
		n, _ := fmt.Scanf("%d", &num)
		if n == 0 {
			break
		}
		for i := int64(2); i <= num; i++ {
			if num == 1 {
				break
			}
			for num%i == 0 {
				fmt.Printf("%d ", i)
				num = num / i
			}
		}
		fmt.Println()
	}
}
