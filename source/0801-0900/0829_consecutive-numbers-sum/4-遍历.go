package main

func main() {

}

// leetcode829_连续整数求和
func consecutiveNumbersSum(N int) int {
	res := 0
	i := 1
	for N > 0 {
		if N%i == 0 {
			res++
		}
		N = N - i
		i++
	}
	return res
}
