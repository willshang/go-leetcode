package main

func main() {

}

// leetcode1745_回文串分割IV
func checkPartitioning(s string) bool {
	for i := 1; i < len(s)-1; i++ {
		a := s[0:i]
		if check(a) == false {
			continue
		}
		for j := len(s) - 1; j > i; j-- {
			c := s[j:]
			if check(c) == false {
				continue
			}
			b := s[i:j]
			if check(b) == true {
				return true
			}
		}
	}
	return false
}

func check(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
