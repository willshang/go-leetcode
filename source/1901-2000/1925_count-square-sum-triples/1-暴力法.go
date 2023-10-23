package main

func main() {

}

// leetcode1925_统计平方和三元组的数目
func countTriples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if i*i+j*j == k*k {
					res++
				}
			}
		}
	}
	return res
}
