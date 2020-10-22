package main

import "fmt"

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00"))
}

func maxScore(s string) int {
	max := 0
	for i := 0; i < len(s)-1; i++ {
		zero := 0
		one := 0
		for j := 0; j <= i; j++ {
			if s[j] == '0' {
				zero++
			}
		}
		for j := i + 1; j < len(s); j++ {
			if s[j] == '1' {
				one++
			}
		}
		if zero+one > max {
			max = zero + one
		}
	}
	return max
}
