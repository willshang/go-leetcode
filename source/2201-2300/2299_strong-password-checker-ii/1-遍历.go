package main

func main() {

}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	s := "!@#$%^&*()-+"
	mm := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		mm[s[i]] = true
	}
	m := make(map[int]int)
	arr := []byte(password)
	for k, v := range arr {
		if k >= 1 {
			if v == arr[k-1] {
				return false
			}
		}
		if '0' <= v && v <= '9' {
			m[1] = 1
		}
		if 'a' <= v && v <= 'z' {
			m[2] = 1
		}
		if 'A' <= v && v <= 'Z' {
			m[3] = 1
		}
		if mm[v] == true {
			m[4] = 1
		}
	}
	return len(m) == 4
}
