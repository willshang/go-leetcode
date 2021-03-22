package main

import "math"

func main() {

}

// leetcode1362_最接近的因数
func closestDivisors(num int) []int {
	res := []int{1, num + 1}
	for i := int(math.Sqrt(float64(num + 2))); i >= 0; i-- {
		if (num+1)%i == 0 {
			return []int{i, (num + 1) / i}
		}
		if (num+2)%i == 0 {
			return []int{i, (num + 2) / i}
		}
	}
	return res
}
