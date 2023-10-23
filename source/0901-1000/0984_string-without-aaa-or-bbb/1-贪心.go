package main

func main() {

}

// leetcode984_不含AAA或BBB的字符串
func strWithout3a3b(a int, b int) string {
	res := make([]byte, 0)
	for a > 0 || b > 0 {
		flagA := false
		if len(res) >= 2 && res[len(res)-1] == res[len(res)-2] {
			if res[len(res)-1] == 'b' {
				flagA = true
			}
		} else if a >= b {
			flagA = true
		}
		if flagA == true {
			res = append(res, 'a')
			a--
		} else {
			res = append(res, 'b')
			b--
		}
	}
	return string(res)
}
