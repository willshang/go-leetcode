package main

func main() {

}

func baseNeg2(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	for n != 0 {
		if n%2 == 0 { // 偶数
			res = "0" + res
			n = n / -2
		} else { // 奇数
			res = "1" + res
			n = (n - 1) / -2
		}
	}
	return res
}
