package main

func main() {

}

// leetcode1006_笨阶乘
func clumsy(N int) int {
	res := 0
	sum := 1
	for i := N; i > 0; i = i - 4 {
		if i == 1 {
			sum = sum * 1
		} else if i == 2 {
			sum = sum * 2 * 1
		} else if i == 3 {
			sum = sum * 3 * 2 / 1
		} else {
			sum = sum*i*(i-1)/(i-2) + (i - 3)
		}
		res = res + sum
		sum = -1
	}
	return res
}
