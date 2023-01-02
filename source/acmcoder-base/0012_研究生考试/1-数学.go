package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)
		sum := a + b + c + d
		if a < 60 || b < 60 || c < 90 || d < 90 || sum < 310 {
			fmt.Println("Fail")
		} else if sum >= 310 && sum <= 349 {
			fmt.Println("Zifei")
		} else {
			fmt.Println("Gongfei")
		}
	}
}
