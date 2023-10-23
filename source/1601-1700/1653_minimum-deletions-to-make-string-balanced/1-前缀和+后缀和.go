package main

func main() {

}

func minimumDeletions(s string) int {
	n := len(s)
	dpA := make([]int, n)
	dpB := make([]int, n)
	if s[0] == 'a' {
		dpA[0] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == 'a' {
			dpA[i] = dpA[i-1] + 1
		} else {
			dpA[i] = dpA[i-1]
		}
	}
	if s[n-1] == 'b' {
		dpB[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == 'b' {
			dpB[i] = dpB[i+1] + 1
		} else {
			dpB[i] = dpB[i+1]
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dpA[i]+dpB[i])
	}
	return n - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
