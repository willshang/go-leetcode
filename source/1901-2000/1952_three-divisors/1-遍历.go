package main

func main() {

}

// leetcode1952_三除数
func isThree(n int) bool {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count == 3
}
