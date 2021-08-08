package main

func main() {

}

func isThree(n int) bool {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count = count + 2
			}
		}
	}
	return count == 3
}
