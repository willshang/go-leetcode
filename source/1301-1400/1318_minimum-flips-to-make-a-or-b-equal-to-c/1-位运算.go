package main

func main() {

}

func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 31; i++ {
		A := (a >> i) & 1
		B := (b >> i) & 1
		C := (c >> i) & 1
		if C == 0 {
			res = res + A + B
		} else {
			if A+B == 0 {
				res = res + 1
			}
		}
	}
	return res
}
