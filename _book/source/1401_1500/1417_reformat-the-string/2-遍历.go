package main

import "fmt"

func main() {
	fmt.Println(reformat("a0b1c2"))
}

// leetcode1417_重新格式化字符串
func reformat(s string) string {
	res := make([]byte, 0)
	m1 := make([]byte, 0)
	m2 := make([]byte, 0)
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			m1 = append(m1, s[i])
		} else {
			m2 = append(m2, s[i])
		}
	}
	if len(m1)-len(m2) == 1 {
		for i := 0; i < len(m2); i++ {
			res = append(res, m1[i])
			res = append(res, m2[i])
		}
		res = append(res, m1[len(m1)-1])
		return string(res)
	} else if len(m2)-len(m1) == 1 {
		for i := 0; i < len(m1); i++ {
			res = append(res, m2[i])
			res = append(res, m1[i])
		}
		res = append(res, m2[len(m2)-1])
		return string(res)
	} else if len(m2) == len(m1) {
		for i := 0; i < len(m1); i++ {
			res = append(res, m2[i])
			res = append(res, m1[i])
		}
		return string(res)
	} else {
		return ""
	}
}
