package main

import "math"

func main() {

}

func closestDivisors(num int) []int {
	res := []int{1, num + 1}
	for i := num + 1; i <= num+2; i++ {
		temp := divide(i)
		if abs(temp[0]-temp[1]) < abs(res[0]-res[1]) {
			res = temp
		}
	}
	return res
}

func divide(n int) []int {
	for i := int(math.Sqrt(float64(n))); i >= 0; i-- {
		if n%i == 0 {
			return []int{i, n / i}
		}
	}
	return nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
