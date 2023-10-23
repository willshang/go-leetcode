package main

import "fmt"

func main() {
	fmt.Println(longestDecomposition("elvtoelvto"))
	fmt.Println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
}

func longestDecomposition(text string) int {
	n := len(text)
	if n <= 1 {
		return n
	}
	res := 0
	left, right := 0, n
	for left < right {
		for k := 1; ; k++ {
			if k > (right-left)/2 { // 长度超过剩下的1/2
				res++
				return res
			}
			if text[left:left+k] == text[right-k:right] {
				res = res + 2 // 可以切割，长度+2
				left = left + k
				right = right - k
				break
			}
		}
	}
	return res
}
