package main

import "fmt"

func main() {
	fmt.Println(ambiguousCoordinates("(123)"))
}

// leetcode816_模糊坐标
func ambiguousCoordinates(s string) []string {
	res := make([]string, 0)
	str := s[1 : len(s)-1]
	n := len(str)
	for i := 1; i <= n-1; i++ { // 枚举左右2边
		left := str[:i]
		right := str[i:]
		a := process(left)
		b := process(right)
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(b); k++ {
				res = append(res, "("+a[j]+", "+b[k]+")")
			}
		}
	}
	return res
}

func process(str string) []string {
	res := make([]string, 0)
	n := len(str)
	for i := 1; i <= n; i++ {
		left := str[:i]
		right := str[i:]
		if judge(left, right) == true {
			if i == n {
				res = append(res, left+""+right)
			} else {
				res = append(res, left+"."+right)
			}
		}
	}
	return res
}

func judge(a, b string) bool {
	if (len(a) >= 2 && a[0] == '0') ||
		(len(b) >= 1 && b[len(b)-1] == '0') {
		return false
	}
	return true
}
