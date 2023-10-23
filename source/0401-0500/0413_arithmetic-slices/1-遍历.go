package main

func main() {

}

// leetcode413_等差数列划分
func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	res := 0
	count := 2
	diff := A[1] - A[0]
	for i := 2; i < n; i++ {
		a := A[i] - A[i-1]
		if a == diff {
			count++
		} else {
			if count >= 3 {
				res = res + getValue(count)
			}
			count = 2
			diff = a
		}
	}
	if count >= 3 {
		res = res + getValue(count)
	}
	return res
}

func getValue(num int) int {
	n := num - 2
	return n * (n + 1) / 2
}
