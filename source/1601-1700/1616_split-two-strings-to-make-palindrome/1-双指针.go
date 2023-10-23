package main

import "fmt"

func main() {
	//fmt.Println(checkPalindromeFormation("ulacfd", "jizalu"))
	//fmt.Println(checkPalindromeFormation("aabcc", "ccdaa"))
	fmt.Println(checkPalindromeFormation("affggddd", "jsssdafa"))
}

func checkPalindromeFormation(a string, b string) bool {
	if judge(a, b) == true || judge(b, a) == true {
		return true
	}
	return false
}

// 判断prefixA+suffixB是否是回文
func judge(a, b string) bool {
	left, right := 0, len(a)-1
	for left < right {
		if a[left] == b[right] {
			left++
			right--
		} else {
			break
		}
	}
	var i, j int
	// 2种判断
	// A1+A2+A3
	// B1+B2+B3
	// 1.reverse(A1)=B3，判断B2是否回文
	i, j = left, right
	for i < j {
		if b[i] == b[j] {
			i++
			j--
		} else {
			break
		}
	}
	if i >= j {
		return true
	}
	// 2.reverse(A1)=B3，判断A2是否回文
	i, j = left, right
	for i < j {
		if a[i] == a[j] {
			i++
			j--
		} else {
			break
		}
	}
	if i >= j {
		return true
	}
	return false
}
