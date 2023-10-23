package main

import "unicode"

func main() {

}

// leetcode2299_强密码检验器II
func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	m := make(map[int]int)
	for i := 0; i < len(password); i++ {
		if i >= 1 {
			if password[i] == password[i-1] {
				return false
			}
		}
		v := rune(password[i])
		switch {
		case unicode.IsLower(v):
			m[1] = 1
		case unicode.IsUpper(v):
			m[2] = 1
		case unicode.IsDigit(v):
			m[3] = 1
		default:
			m[4] = 1
		}
	}
	return len(m) == 4
}
