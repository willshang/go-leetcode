package main

func main() {

}

func minFlipsMonoIncr(S string) int {
	n := len(S)
	a := 0 // 0 结尾
	b := 0 // 1 结尾
	if S[0] == '1' {
		a = 1
	} else {
		b = 1
	}
	for i := 1; i < n; i++ {
		if S[i] == '1' {
			a, b = a+1, min(a, b)
		} else {
			a, b = a, min(a, b)+1
		}
	}
	return min(a, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
