package main

import "math"

func main() {

}

func isThree(n int) bool {
	target := int(math.Sqrt(float64(n)))
	return target*target == n && isPrime(target) == true
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}
