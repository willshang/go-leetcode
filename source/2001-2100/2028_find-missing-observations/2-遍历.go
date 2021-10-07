package main

func main() {

}

func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	m := len(rolls)
	for i := 0; i < m; i++ {
		sum = sum + rolls[i]
	}
	total := mean * (n + m)
	left := total - sum
	if left > n*6 || left < n {
		return nil
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	left = left - n
	for i := 0; i < n; i++ {
		if left < 6 {
			res[i] = res[i] + left
			break
		}
		res[i] = res[i] + 5
		left = left - 5
	}
	return res
}
