package main

import (
	"math"
)

func main() {

}

func divisorSubstrings(num int, k int) int {
	res := 0
	target := int(math.Pow10(k))
	for v := num; v >= target/10; v = v / 10 {
		value := v % target
		if value > 0 && num%value == 0 {
			res++
		}
	}
	return res
}
