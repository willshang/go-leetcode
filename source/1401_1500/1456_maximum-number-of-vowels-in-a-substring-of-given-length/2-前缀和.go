package main

import "fmt"

func main() {
	fmt.Println(maxVowels("ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33))
}

func maxVowels(s string, k int) int {
	res := 0
	total := 0
	arr := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) == true {
			total++
		}
		arr[i+1] = total
		if i >= k-1 {
			res = max(res, arr[i+1]-arr[i-k+1])
		}
	}
	return res
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' ||
		b == 'i' || b == 'o' || b == 'u'
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
