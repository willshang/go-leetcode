package main

import "fmt"

func main() {
	fmt.Println(lexicalOrder(200))
}

func lexicalOrder(n int) []int {
	res := make([]int, 0)
	num := 1
	for {
		if num <= n {
			res = append(res, num)
			num = num * 10
		} else {
			num = num / 10
			for num%10 == 9 {
				num = num / 10
			}
			if num == 0 {
				break
			}
			num++
		}
	}
	return res
}
