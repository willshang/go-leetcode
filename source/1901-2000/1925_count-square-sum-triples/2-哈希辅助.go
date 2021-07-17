package main

func main() {

}

func countTriples(n int) int {
	res := 0
	m := make(map[int]bool)
	for i := 1; i <= n; i++ {
		m[i*i] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if m[i*i+j*j] == true {
				res++
			}
		}
	}
	return res
}
