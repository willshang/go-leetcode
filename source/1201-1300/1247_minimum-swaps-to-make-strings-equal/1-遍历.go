package main

func main() {

}

// leetcode1247_交换字符使得字符串相同
func minimumSwap(s1 string, s2 string) int {
	a, b := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] { // 相同不需要交换
			continue
		}
		if s1[i] == 'x' {
			a++
		} else {
			b++
		}
	}
	// a => x y
	// b => y x
	if a%2+b%2 == 1 {
		return -1
	}
	return a/2 + b/2 + 2*(a%2)
}
