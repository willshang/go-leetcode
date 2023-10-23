package main

func main() {

}

func minFlipsMonoIncr(S string) int {
	n := len(S)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1]
		if S[i-1] == '1' {
			arr[i]++
		}
	}
	res := n
	for i := 0; i <= n; i++ {
		left := arr[i]
		right := n - i - (arr[n] - arr[i])
		res = min(res, left+right)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
