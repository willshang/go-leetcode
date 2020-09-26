package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
}

// leetcode1071_字符串的最大公因子
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	if str1 > str2 {
		str1, str2 = str2, str1
	}
	return str1[:gcd(len(str2), len(str1))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
