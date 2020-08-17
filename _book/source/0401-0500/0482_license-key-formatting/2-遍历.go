package main

import (
	"fmt"
)

func main() {
	S := "5F3Z-2e-9-w"
	K := 4
	fmt.Println(licenseKeyFormatting(S, K))

	S = "2-5g-3-J"
	K = 2
	fmt.Println(licenseKeyFormatting(S, K))
}

// leetcode482_密钥格式化
func licenseKeyFormatting(S string, K int) string {
	res := make([]rune, 0)
	temp := []rune(S)
	count := 0
	for i := len(temp) - 1; i >= 0; i-- {
		value := temp[i]
		if value >= 'a' {
			value = value - 'a' + 'A'
		}
		if value == '-' {
			continue
		}
		count++
		res = append([]rune{value}, res...)
		if count == K {
			res = append([]rune{'-'}, res...)
			count = 0
		}
	}
	if len(res) == 0 {
		return ""
	}
	if res[0] == '-' {
		res = res[1:]
	}
	return string(res)
}
