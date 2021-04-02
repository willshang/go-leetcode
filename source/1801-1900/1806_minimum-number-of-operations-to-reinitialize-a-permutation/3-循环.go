package main

func main() {

}

func reinitializePermutation(n int) int {
	res := 0
	target := 1
	// 反向思路，只考虑1的变换
	for {
		if target*2 < n {
			target = target * 2
		} else {
			target = target*2 + 1 - n
		}
		res++
		if target == 1 {
			break
		}
	}
	return res
}
