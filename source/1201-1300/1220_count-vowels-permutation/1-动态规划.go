package main

func main() {

}

// leetcode1220_统计元音字母序列的数目
var mod = 1000000007

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for k := 2; k <= n; k++ {
		tempA := (e + i + u) % mod
		tempE := (a + i) % mod
		tempI := (e + o) % mod
		tempO := i % mod
		tempU := (i + o) % mod
		a, e, i, o, u = tempA, tempE, tempI, tempO, tempU
	}
	return (a + e + i + o + u) % mod
}
