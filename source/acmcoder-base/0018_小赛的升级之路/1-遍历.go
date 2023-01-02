package main

import "fmt"

func main() {
	var n, a int
	for {
		v, _ := fmt.Scan(&n, &a)
		if v == 0 {
			break
		} else {
			for i := 0; i < n; i++ {
				var temp int
				fmt.Scan(&temp)
				if a >= temp {
					a = a + temp
				} else {
					a = a + gcd(a, temp)
				}
			}
			fmt.Println(a)
		}
	}
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(a%b, b)
}
