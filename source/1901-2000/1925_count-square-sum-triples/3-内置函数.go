package main

import "math"

func main() {

}

func countTriples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			k := int(math.Sqrt(float64(i*i + j*j)))
			if k*k == i*i+j*j && k <= n {
				res++
			}
		}
	}
	return res
}
