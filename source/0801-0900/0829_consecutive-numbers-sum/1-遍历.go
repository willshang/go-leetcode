package main

func main() {

}

func consecutiveNumbersSum(N int) int {
	res := 1
	sum := 0
	for i := 1; i < N; i++ {
		sum = sum + i // 1~i
		left := N - sum
		if left > 0 {
			if left%(i+1) == 0 { // 划分为i+1个数
				res++
			}
		} else {
			break
		}
	}
	return res
}
