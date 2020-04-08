package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(27))
}

func convertToTitle(n int) string {
	if n <= 26 {
		return string('A' + n - 1)
	}
	y := n % 26
	if y == 0 {
		// 26的倍数 如52%26=0 => AZ
		return convertToTitle((n-y-1)/26) + convertToTitle(26)
	}
	return convertToTitle((n-y)/26) + convertToTitle(y)
}
