package main

func main() {

}

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	// f(2k)/f(2k+1) = 2(k+1-f(k))，最后奇数2k+1不影响结果
	return 2 * (n/2 + 1 - lastRemaining(n/2))
}
