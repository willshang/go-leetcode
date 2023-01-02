package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	res := getResult(a, b)
	fmt.Printf("%.2f%%\n", res)
}

func getResult(a, b string) float64 {
	n := len(a)
	count := 0
	for i := 0; i < n; i++ {
		c := 0
		d := int(b[i] - '0')
		if ('0' <= a[i] && a[i] <= '9') ||
			('a' <= a[i] && a[i] <= 'z') ||
			('A' <= a[i] && a[i] <= 'Z') {
			c = 1
		}
		if c == d {
			count++
		}
	}
	return float64(count) / float64(n) * 100
}
