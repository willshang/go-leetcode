package main

import "fmt"

func main() {

}

func minBitFlips(start int, goal int) int {
	res := 0
	a := fmt.Sprintf("%032b", start)
	b := fmt.Sprintf("%032b", goal)
	for i := 0; i < 32; i++ {
		if a[i] != b[i] {
			res++
		}
	}
	return res
}
