package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b, c int
		_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
		var result int
		if a > b {
			result = a - c + b
		} else {
			result = b - c + a
		}
		if result > 0 {
			fmt.Print("Case #", i+1, ": ", true, "\n")
		} else {
			fmt.Print("Case #", i+1, ": ", false, "\n")
		}
	}
}
